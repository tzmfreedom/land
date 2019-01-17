package builtin

import (
	"github.com/tzmfreedom/goland/ast"
)

var ListType = createListType()

func CreateListType(classType *ast.ClassType) *ast.ClassType {
	return &ast.ClassType{
		Name:            "List",
		Modifiers:       ListType.Modifiers,
		Constructors:    ListType.Constructors,
		InstanceFields:  ListType.InstanceFields,
		InstanceMethods: ListType.InstanceMethods,
		StaticFields:    ListType.StaticFields,
		StaticMethods:   ListType.StaticMethods,
		Extra: map[string]interface{}{
			"generics": []*ast.ClassType{classType},
		},
	}
}

func CreateListTypeRef(typeRef *ast.TypeRef) *ast.TypeRef {
	return &ast.TypeRef{
		Name:       []string{"List"},
		Parameters: []ast.Node{typeRef},
	}
}

func CreateListTypeParameter(typeRef *ast.TypeRef) *ast.Parameter {
	return &ast.Parameter{
		TypeRef: CreateListTypeRef(typeRef),
		Name:    "_",
	}
}

func CreateListObject(classType *ast.ClassType, records []*ast.Object) *ast.Object {
	listObj := ast.CreateObject(ListType)
	listObj.GenericType = []*ast.ClassType{classType}
	listObj.Extra["records"] = records
	return listObj
}

var t1Parameter = &ast.Parameter{
	TypeRef: &ast.TypeRef{
		Name:       []string{"T:1"},
		Parameters: []ast.Node{},
	},
	Name: "_",
}

var t2Parameter = &ast.Parameter{
	TypeRef: &ast.TypeRef{
		Name:       []string{"T:2"},
		Parameters: []ast.Node{},
	},
	Name: "_",
}

var t1TypeRef = &ast.TypeRef{
	Name:       []string{"T:1"},
	Parameters: []ast.Node{},
}

var t2TypeRef = &ast.TypeRef{
	Name:       []string{"T:2"},
	Parameters: []ast.Node{},
}

func createListType() *ast.ClassType {
	instanceMethods := ast.NewMethodMap()
	instanceMethods.Set(
		"add",
		[]*ast.Method{
			ast.CreateMethod(
				"add",
				nil,
				[]*ast.Parameter{t1Parameter},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					records := this.Extra["records"].([]*ast.Object)
					listElement := params[0]
					this.Extra["records"] = append(records, listElement)
					return nil
				},
			),
		},
	)
	instanceMethods.Set(
		"size",
		[]*ast.Method{
			ast.CreateMethod(
				"size",
				integerTypeRef,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					return NewInteger(len(this.Extra["records"].([]*ast.Object)))
				},
			),
		},
	)

	classType := ast.CreateClass(
		"List",
		[]*ast.Method{
			{
				Modifiers:  []*ast.Modifier{ast.PublicModifier()},
				Parameters: []*ast.Parameter{},
				NativeFunction: func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					return nil
				},
			},
			{
				Modifiers: []*ast.Modifier{ast.PublicModifier()},
				Parameters: []*ast.Parameter{
					{
						TypeRef: &ast.TypeRef{
							Name: []string{"List"},
							Parameters: []ast.Node{
								&ast.TypeRef{
									Name: []string{"T:1"},
								},
							},
						},
						Name: "list",
					},
				},
				NativeFunction: func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					listObj := params[0]
					listParams := listObj.Extra["records"].([]*ast.Object)
					newListParams := make([]*ast.Object, len(listParams))
					for i, listParam := range listParams {
						newListParams[i] = &ast.Object{
							ClassType: listParam.ClassType,
						}
					}
					this.Extra = map[string]interface{}{
						"records": newListParams,
					}
					return nil
				},
			},
		},
		instanceMethods,
		nil,
	)
	return classType
}

func init() {
	primitiveClassMap.Set("list", ListType)
}
