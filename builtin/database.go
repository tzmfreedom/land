package builtin

import (
	"github.com/tzmfreedom/go-soapforce"
	"github.com/tzmfreedom/goland/ast"
)

func init() {
	staticMethods := NewMethodMap()
	method := CreateMethod(
		"insert",
		"", // TODO: implement
		func(this interface{}, params []interface{}, options ...interface{}) interface{} {
			sobj := params[0].(*Object)
			record := &soapforce.SObject{}
			for k, v := range sobj.InstanceFields.All() {
				record.Fields[k] = v.Value()
			}
			client := NewSoapClient()
			client.Create([]*soapforce.SObject{record})
			return Null
		},
	)
	staticMethods.Set("insert", []ast.Node{method})

	databaseClass := CreateClass(
		"Database",
		[]*ast.ConstructorDeclaration{},
		nil,
		staticMethods,
	)
	primitiveClassMap.Set("list", databaseClass)
}
