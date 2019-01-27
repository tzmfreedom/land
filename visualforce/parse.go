package main

import (
	"bytes"
	"encoding/xml"
	"io/ioutil"
	"fmt"
	"strings"
	"github.com/tzmfreedom/goland/ast"
	"regexp"
	"github.com/tzmfreedom/goland/builtin"
)

type Env struct {
	Fields *ast.ObjectMap
	Methods map[string]struct{}
}

func NewEnv() *Env {
	return &Env{
		Fields: ast.NewObjectMap(),
		Methods: map[string]struct{}{},
	}
}

type Node struct {
	XMLName xml.Name
	Attrs   []xml.Attr `xml:"-"`
	Content []byte     `xml:",innerxml"`
	Nodes   []Node     `xml:",any"`
}

func (n *Node) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	n.Attrs = start.Attr
	type node Node

	return d.DecodeElement((*node)(n), &start)
}

func Walk(nodes []Node, e *Env) string {
	parts := make([]string, len(nodes))
	for i, n := range nodes {
		parts[i] = render(n, e)
	}
	return strings.Join(parts, "\n")
}

func render(n Node, e *Env) string {
	var part string
	if n.XMLName.Space == "" {
		part = renderNode(n)
	} else if n.XMLName.Space == "apex" {
		switch n.XMLName.Local {
		case "page":
			attributeValue(n, "controller")
			// evaluate attribute
			env := NewEnv()
			env.Fields.Set("accName", builtin.NewString("hoge"))
			body := Walk(n.Nodes, env)
			part = fmt.Sprintf(`<html>
<head></head>
<body>
%s
</body>
</html>`, body)
		case "form":
			body := Walk(n.Nodes, e)
			part = fmt.Sprintf(`<form>
%s
</form>`, body)
		case "outputLabel":
			value := attributeValue(n, "value")
			_, valueObj := bindValue(value, e)
			part = fmt.Sprintf(`<label>%s</label>`, builtin.String(valueObj))
		case "inputField":
			attrValue := attributeValue(n, "value")
			name, valueObj := bindValue(attrValue, e)
			value := builtin.String(valueObj)
			switch valueObj.ClassType {
			case builtin.IntegerType:
				part = fmt.Sprintf(`<input type="text" name="%s" class="text-field" value="%s" />`, name, value)
			case builtin.StringType:
				part = fmt.Sprintf(`<input type="text" name="%s" class="text-field" value="%s" />`, name, value)
			case builtin.DoubleType:
				part = fmt.Sprintf(`<input type="text" name="%s" class="text-field" value="%s" />`, name, value)
			case builtin.BooleanType:
				part = fmt.Sprintf(`<input type="checkbox" name="%s" class="checkbox-field" value="%s" />`, name, value)
			case builtin.DateType:
				part = fmt.Sprintf(`<input type="text" name="%s" class="date-field" value="%s" />`, name, value)
			}
		}
	}
	return part
}

func main() {
	data, err := ioutil.ReadFile("sample.page")
	if err != nil {
		panic(err)
	}
	buf := bytes.NewBuffer(data)
	dec := xml.NewDecoder(buf)

	var n Node
	err = dec.Decode(&n)
	if err != nil {
		panic(err)
	}
	body := Walk([]Node{n}, nil)
	fmt.Println(body)
}

func renderNode(n Node) string {
	b, err := xml.Marshal(n)
	if err != nil {
		panic(err)
	}
	return string(b)
}

func bindValue(value string, e *Env) (string, *ast.Object) {
	r := regexp.MustCompile(`{!([^{}]+)}`)
	sub := r.FindStringSubmatch(value)
	if len(sub) == 0 {
		return "", builtin.NewString(value)
	}
	name := sub[1]
	v, ok := e.Fields.Get(name)
	if !ok {
		panic("not found: " + name)
	}
	return name, v
}

func attributeValue(n Node, name string) string {
	for _, attr := range n.Attrs {
		if attr.Name.Local == name {
			return attr.Value
		}
	}
	return ""
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

func handleFunc() {

}