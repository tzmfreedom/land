package visualforce

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"regexp"
	"strings"

	"net/http"

	"html/template"

	"github.com/tzmfreedom/goland/ast"
	"github.com/tzmfreedom/goland/builtin"
	"github.com/tzmfreedom/goland/compiler"
	"github.com/tzmfreedom/goland/interpreter"
)

type Node struct {
	XMLName xml.Name
	Attrs   []xml.Attr `xml:"-"`
	Content []byte     `xml:",innerxml"`
	Nodes   []Node     `xml:",any"`
}

type Attribute struct {
	Data map[string]string
}

func (attr *Attribute) Set(k string, v string) {
	attr.Data[strings.ToLower(k)] = v
}

func (attr *Attribute) Get(k string) string {
	return attr.Data[strings.ToLower(k)]
}

func (n *Node) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	n.Attrs = start.Attr
	type node Node

	return d.DecodeElement((*node)(n), &start)
}

var renderFunction = map[string]func(n Node, c *ast.Object) string{}

var pageTmpl *template.Template
var formTmpl *template.Template
var inputFieldTmpl *template.Template
var labelTmpl *template.Template

type PageParameter struct {
	ShowHeader bool
	Body       string
}

type LabelParameter struct {
	Value string
}

type InputFieldParameter struct {
	Type  string
	Name  string
	Value string
}

type FormParameter struct {
	Body string
}

func handleRequest(i *interpreter.Interpreter, r *http.Request, w http.ResponseWriter) {
	// initialize page to specify controller
	pagePath := r.URL.Path[1:]
	n, err := createNode(pagePath)
	if err != nil {
		w.WriteHeader(400)
		return
	}
	attrs := attributeValues(n)
	controller := attrs.Get("controller")
	r.ParseForm()
	method := r.Form["__action"][0]
	if method == "" {
		panic("method not blank")
	}
	state := map[string]*ast.Object{}
	// run action method
	c, retValue, err := i.BindAndRun(controller, method, r.Form, state)
	if err != nil {
		panic(err)
	}
	// evaluate return value
	if retValue == builtin.Null {
		pagePath = r.URL.Path[1:]
	} else {
		pagePath = retValue.Extra["url"].(*ast.Object).StringValue()
		if pagePath != r.Referer() {
			http.Redirect(w, r, pagePath, http.StatusFound)
			return
		}
	}
	// render page if return value is same page
	for k, v := range c.InstanceFields.All() {
		state[k] = v
	}
	n, err = createNode(pagePath)
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		return
	}
	attrs = attributeValues(n)
	body := renderNodes(n.Nodes, c)
	pageTmpl.Execute(w, PageParameter{
		Body:       body,
		ShowHeader: attrs.Get("showHeader") != "false",
	})
}

func renderPage(n Node, i *interpreter.Interpreter) (string, error) {
	if n.XMLName.Space != "apex" || n.XMLName.Local != "page" {
		panic("root tag must be apex:page")
	}
	attrs := attributeValues(n)
	// evaluate attribute
	controller, _, err := i.BindAndRun(attrs.Get("controller"), "", nil, nil)
	if err != nil {
		return "", err
	}
	body := renderNodes(n.Nodes, controller)
	return renderTemplate(pageTmpl, PageParameter{
		Body:       body,
		ShowHeader: attrs.Get("showHeader") != "false",
	}), nil
}

func renderNodes(nodes []Node, c *ast.Object) string {
	parts := make([]string, len(nodes))
	for i, n := range nodes {
		parts[i] = renderNode(n, c)
	}
	return strings.Join(parts, "\n")
}

func renderNode(n Node, c *ast.Object) string {
	if n.XMLName.Space == "apex" {
		return renderFunction[n.XMLName.Local](n, c)
	}
	return renderHtmlTag(n)
}

func renderHtmlTag(n Node) string {
	b, err := xml.Marshal(n)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func getClassType(controller *ast.Object, value string) *ast.ClassType {
	r := regexp.MustCompile(`{!([^{}]+)}`)
	sub := r.FindStringSubmatch(value)
	if len(sub) == 0 {
		return builtin.StringType
	}
	names := strings.Split(sub[1], ".")
	f, err := compiler.FindInstanceField(controller.ClassType, names[0], compiler.MODIFIER_PUBLIC_ONLY, false)
	for _, name := range names[1:] {
		f, err = compiler.FindInstanceField(f.Type, name, compiler.MODIFIER_PUBLIC_ONLY, false)
		if err != nil {
			panic("not found: " + name)
		}
	}
	return f.Type
}

func bindMethod(value string, c *ast.Object) (string, *ast.Method) {
	r := regexp.MustCompile(`{!([^{}]+)}`)
	sub := r.FindStringSubmatch(value)
	if len(sub) == 0 {
		panic("must bind {!xxx}")
	}
	name := sub[1]
	names := strings.Split(name, ".")
	var receiver = c
	var ok bool
	for _, n := range names[:len(names)-1] {
		receiver, ok = receiver.InstanceFields.Get(n)
		if !ok {
			panic("not found: " + n)
		}
	}
	methodName := names[len(names)-1]
	_, method, err := interpreter.FindInstanceMethod(receiver, methodName, []*ast.Object{}, compiler.MODIFIER_ALL_OK)
	if err != nil {
		panic("not found: " + methodName)
	}
	return methodName, method
}

func bindInstanceField(value string, c *ast.Object) (string, *ast.Object) {
	r := regexp.MustCompile(`{!([^{}]+)}`)
	sub := r.FindStringSubmatch(value)
	if len(sub) == 0 {
		return "", builtin.NewString(value)
	}
	name := sub[1]
	names := strings.Split(name, ".")
	var receiver = c
	var ok bool
	for _, n := range names {
		receiver, ok = receiver.InstanceFields.Get(n)
		if !ok {
			panic("not found: " + n)
		}
	}
	return name, receiver
}

func attributeValues(n Node) *Attribute {
	attrs := &Attribute{
		Data: map[string]string{},
	}
	for _, attr := range n.Attrs {
		attrs.Set(attr.Name.Local, attr.Value)
	}
	return attrs
}

func childNodes(n Node, space, local string) []Node {
	childNodes := []Node{}
	for _, node := range n.Nodes {
		if node.XMLName.Space == space && node.XMLName.Local == local {
			childNodes = append(childNodes, node)
		}
	}
	return childNodes
}

func createNode(pagePath string) (Node, error) {
	data, err := ioutil.ReadFile(pagePath)
	if err != nil {
		return Node{}, err
	}
	buf := bytes.NewBuffer(data)
	dec := xml.NewDecoder(buf)

	var n Node
	err = dec.Decode(&n)
	if err != nil {
		panic(err)
	}
	return n, nil
}

func Render(pagePath string, i *interpreter.Interpreter) (string, error) {
	n, err := createNode(pagePath)
	if err != nil {
		return "", err
	}
	return renderPage(n, i)
}

func Server(i *interpreter.Interpreter) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		switch r.Method {
		case http.MethodGet:
			body, err := Render(r.URL.Path[1:], i)
			if err != nil {
				w.WriteHeader(404)
				return
			}
			fmt.Fprint(w, body)
		case http.MethodPost:
			handleRequest(i, r, w)
		}
	})
	fmt.Println("listening to 0.0.0.0:8080")
	http.ListenAndServe(":8080", nil)
}

func init() {
	renderFunction["form"] = func(n Node, c *ast.Object) string {
		body := renderNodes(n.Nodes, c)
		return renderTemplate(formTmpl, FormParameter{
			Body: body,
		})
	}
	renderFunction["outputLabel"] = func(n Node, c *ast.Object) string {
		attr := attributeValues(n)
		_, valueObj := bindInstanceField(attr.Get("value"), c)
		return renderTemplate(labelTmpl, LabelParameter{
			Value: builtin.String(valueObj),
		})
	}
	renderFunction["commandButton"] = func(n Node, c *ast.Object) string {
		attr := attributeValues(n)
		attrValue := attr.Get("value")
		name, valueObj := bindInstanceField(attrValue, c)
		value := builtin.String(valueObj)

		attrAction := attr.Get("action")
		actionName, _ := bindMethod(attrAction, c)

		return fmt.Sprintf(`<input type="hidden" name="__action" value="%s" />
<input type="submit" name="%s" value="%s" />`, actionName, name, value)
	}
	renderFunction["inputField"] = func(n Node, c *ast.Object) string {
		attr := attributeValues(n)
		attrValue := attr.Get("value")
		name, valueObj := bindInstanceField(attrValue, c)
		value := builtin.String(valueObj)
		classType := getClassType(c, attrValue)
		return renderTemplate(inputFieldTmpl, InputFieldParameter{
			Type:  classType.Name,
			Name:  name,
			Value: value,
		})
	}
}

func renderTemplate(tmpl *template.Template, param interface{}) string {
	buf := new(bytes.Buffer)
	tmpl.Execute(buf, param)
	return buf.String()
}

var funcMap = template.FuncMap{
	"safehtml": func(text string) template.HTML { return template.HTML(text) },
}

func createTemplate(name string) *template.Template {
	pagePath := "./visualforce/templates/" + name + ".html"
	content, err := ioutil.ReadFile(pagePath)
	if err != nil {
		panic(err)
	}
	tmpl, err := template.New(name).Funcs(funcMap).Parse(string(content))
	if err != nil {
		panic(err)
	}
	return tmpl
}

func init() {
	pageTmpl = createTemplate("page")
	formTmpl = createTemplate("form")
	inputFieldTmpl = createTemplate("inputField")
	labelTmpl = createTemplate("label")
}
