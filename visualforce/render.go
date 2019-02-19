package visualforce

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"regexp"
	"strings"

	"html/template"

	"encoding/base64"
	"encoding/json"

	"github.com/tzmfreedom/land/ast"
	"github.com/tzmfreedom/land/builtin"
	"github.com/tzmfreedom/land/compiler"
	"github.com/tzmfreedom/land/interpreter"
)

var renderFunction = map[string]func(n Node, c *ast.Object) string{}

var templateStore map[string]*template.Template

type PageParameter struct {
	ShowHeader bool
	Body       string
}

type LabelParameter struct {
	Value string
}

type InputFieldParameter struct {
	Type  string
	Label string
	Name  string
	Value string
}

type FormParameter struct {
	Body         string
	ViewState    string
	RawViewState string
}

type PageBlockParameter struct {
	Body string
}

type PageBlockSectionParameter struct {
	Title string
	Body  string
}

type CommandButtonParameter struct {
	Action string
	Name   string
	Value  string
}

func renderPage(n Node, i *interpreter.Interpreter) (string, error) {
	if n.XMLName.Space != "apex" || n.XMLName.Local != "page" {
		panic("root tag must be apex:page")
	}
	attrs := n.attributeValues()
	// evaluate attribute
	controller, _, err := i.BindAndRun(attrs.Get("controller"), "", nil, nil)
	if err != nil {
		return "", err
	}
	body := renderNodes(n.Nodes, controller)
	return renderTemplate("page", PageParameter{
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

func render(pagePath string, i *interpreter.Interpreter) (string, error) {
	n, err := createNode(pagePath)
	if err != nil {
		return "", err
	}
	return renderPage(n, i)
}

func renderTemplate(templateName string, param interface{}) string {
	buf := new(bytes.Buffer)
	tmpl := templateStore[templateName]
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

func serializeViewState(o *ast.Object) interface{} {
	switch o.ClassType {
	case builtin.StringType, builtin.IntegerType, builtin.DoubleType, builtin.BooleanType:
		return o.Value()
	case builtin.NullType:
		return nil
	}
	data := map[string]interface{}{}
	for k, v := range o.InstanceFields.All() {
		data[k] = serializeViewState(v)
	}
	return data
}

func init() {
	renderFunction["form"] = func(n Node, c *ast.Object) string {
		body := renderNodes(n.Nodes, c)
		viewstateObj := serializeViewState(c).(map[string]interface{})
		viewstate, err := json.Marshal(viewstateObj)
		if err != nil {
			panic(err)
		}
		b64viewstate := base64.StdEncoding.EncodeToString(viewstate)
		return renderTemplate("form", FormParameter{
			Body:         body,
			RawViewState: string(viewstate),
			ViewState:    b64viewstate,
		})
	}
	renderFunction["outputLabel"] = func(n Node, c *ast.Object) string {
		attr := n.attributeValues()
		_, valueObj := bindInstanceField(attr.Get("value"), c)
		return renderTemplate("label", LabelParameter{
			Value: builtin.String(valueObj),
		})
	}
	renderFunction["commandButton"] = func(n Node, c *ast.Object) string {
		attr := n.attributeValues()
		attrValue := attr.Get("value")
		name, valueObj := bindInstanceField(attrValue, c)
		value := builtin.String(valueObj)

		attrAction := attr.Get("action")
		actionName, _ := bindMethod(attrAction, c)
		return renderTemplate("commandButton", CommandButtonParameter{
			Action: actionName,
			Name:   name,
			Value:  value,
		})
	}
	renderFunction["inputField"] = func(n Node, c *ast.Object) string {
		attr := n.attributeValues()
		attrValue := attr.Get("value")
		name, valueObj := bindInstanceField(attrValue, c)

		// TODO: impl
		names := strings.Split(name, ".")
		value := builtin.String(valueObj)
		classType := getClassType(c, attrValue)
		return renderTemplate("inputField", InputFieldParameter{
			Type:  classType.Name,
			Label: names[len(names)-1],
			Name:  name,
			Value: value,
		})
	}
	renderFunction["pageBlock"] = func(n Node, c *ast.Object) string {
		body := renderNodes(n.Nodes, c)
		return renderTemplate("pageBlock", PageBlockParameter{
			Body: body,
		})
	}
	renderFunction["pageBlockSection"] = func(n Node, c *ast.Object) string {
		attr := n.attributeValues()
		attrValue := attr.Get("title")
		title, _ := bindInstanceField(attrValue, c)
		body := renderNodes(n.Nodes, c)
		return renderTemplate("pageBlockSection", PageBlockSectionParameter{
			Title: title,
			Body:  body,
		})
	}

	vfTags := []string{
		"page",
		"pageBlock",
		"pageBlockSection",
		"commandButton",
		"form",
		"inputField",
		"label",
	}
	templateStore = map[string]*template.Template{}
	for _, vgTag := range vfTags {
		templateStore[vgTag] = createTemplate(vgTag)
	}
}
