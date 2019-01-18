package builtin

import (
	"github.com/tzmfreedom/goland/ast"
)

var ListType = &ast.ClassType{Name: "List"}

func CreateListType(classType *ast.ClassType) *ast.ClassType {
	return &ast.ClassType{
		Name:            "List",
		Modifiers:       ListType.Modifiers,
		Constructors:    ListType.Constructors,
		InstanceFields:  ListType.InstanceFields,
		InstanceMethods: ListType.InstanceMethods,
		StaticFields:    ListType.StaticFields,
		StaticMethods:   ListType.StaticMethods,
		Generics:        []*ast.ClassType{classType},
	}
}

func CreateListTypeRef(typeRef *ast.TypeRef) *ast.TypeRef {
	return &ast.TypeRef{
		Name:       []string{"List"},
		Parameters: []*ast.TypeRef{typeRef},
	}
}

func CreateListTypeParameter(classType *ast.ClassType) *ast.Parameter {
	return &ast.Parameter{
		Type: CreateListType(classType),
		Name: "_",
	}
}

func CreateListObject(classType *ast.ClassType, records []*ast.Object) *ast.Object {
	listObj := ast.CreateObject(ListType)
	listObj.GenericType = []*ast.ClassType{classType}
	listObj.Extra["records"] = records
	return listObj
}

var T1type = &ast.ClassType{Name: "T:1"}
var T2type = &ast.ClassType{Name: "T:2"}

var t1Parameter = &ast.Parameter{
	Type: T1type,
	Name: "_",
}

var t2Parameter = &ast.Parameter{
	Type: T2type,
	Name: "_",
}

var t1TypeRef = &ast.TypeRef{
	Name:       []string{"T:1"},
	Parameters: []*ast.TypeRef{},
}

var t2TypeRef = &ast.TypeRef{
	Name:       []string{"T:2"},
	Parameters: []*ast.TypeRef{},
}

func createListType() {
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
				IntegerType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					return NewInteger(len(this.Extra["records"].([]*ast.Object)))
				},
			),
		},
	)

	ListType.Constructors = []*ast.Method{
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
					Type: CreateListType(T1type),
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
	}
	ListType.InstanceFields = ast.NewFieldMap()
	ListType.StaticFields = ast.NewFieldMap()
	ListType.InstanceMethods = instanceMethods
	ListType.StaticMethods = ast.NewMethodMap()
}

func init() {
	createListType()
	primitiveClassMap.Set("list", ListType)
}
