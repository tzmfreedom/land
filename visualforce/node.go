package visualforce

import (
	"encoding/xml"
	"strings"

	"bytes"
	"io/ioutil"
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

func (n *Node) childNodes(space, local string) []Node {
	childNodes := []Node{}
	for _, node := range n.Nodes {
		if node.XMLName.Space == space && node.XMLName.Local == local {
			childNodes = append(childNodes, node)
		}
	}
	return childNodes
}

func (n *Node) attributeValues() *Attribute {
	attrs := &Attribute{
		Data: map[string]string{},
	}
	for _, attr := range n.Attrs {
		attrs.Set(attr.Name.Local, attr.Value)
	}
	return attrs
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
