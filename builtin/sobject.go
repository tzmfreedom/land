package builtin

import (
	"github.com/tzmfreedom/goland/ast"
)

var account = &ClassType{
	Name: "Account",
	InstanceFields: &FieldMap{
		Data: map[string]*Field{
			"name": {
				Type: &ast.TypeRef{Name: []string{"String"}},
				Name: "name",
				Modifiers: []ast.Node{
					&ast.Modifier{
						Name: "public",
					},
				},
			},
			"website": {
				Type: &ast.TypeRef{Name: []string{"String"}},
				Name: "website",
				Modifiers: []ast.Node{
					&ast.Modifier{
						Name: "public",
					},
				},
			},
			"field1__c": {
				Type: &ast.TypeRef{Name: []string{"String"}},
				Name: "field1__c",
				Modifiers: []ast.Node{
					&ast.Modifier{
						Name: "public",
					},
				},
			},
		},
	},
}

func init() {
	primitiveClassMap.Set("account", account)
}
