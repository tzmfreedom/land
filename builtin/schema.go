package builtin

import "github.com/tzmfreedom/land/ast"

var schemaSObjectType *ast.ClassType
var describeSObjectResultType *ast.ClassType

func init() {
	schema := ast.CreateClass(
		"Schema",
		nil,
		nil,
		&ast.MethodMap{
			Data: map[string][]*ast.Method{
				"getGlobalDescribe": {
					&ast.Method{
						Name:       "getGlobalDescribe",
						Modifiers:  []*ast.Modifier{ast.PublicModifier()},
						Parameters: []*ast.Parameter{},
						ReturnType: CreateMapType(StringType, schemaSObjectType),
						NativeFunction: func(this *ast.Object, parameter []*ast.Object, extra map[string]interface{}) interface{} {
							mapType := CreateMapType(StringType, schemaSObjectType)
							newObj := ast.CreateObject(mapType)
							values := map[string]*ast.Object{}
							for name, _ := range sObjects {
								valueObj := ast.CreateObject(schemaSObjectType)
								valueObj.Extra["type"] = name
								values[name] = valueObj
							}
							newObj.Extra["values"] = values
							return newObj
						},
					},
				},
			},
		},
	)

	primitiveClassMap.Set("Schema", schema)

	classMap := ast.NewClassMap()

	schemaSObjectType = ast.CreateClass(
		"SObjectType",
		nil,
		nil,
		&ast.MethodMap{
			Data: map[string][]*ast.Method{
				"getDescribe": {
					&ast.Method{
						Name:       "getDescribe",
						Modifiers:  []*ast.Modifier{ast.PublicModifier()},
						Parameters: []*ast.Parameter{},
						ReturnType: describeSObjectResultType,
						NativeFunction: func(this *ast.Object, parameter []*ast.Object, extra map[string]interface{}) interface{} {
							sObj := sObjects[this.Extra["type"].(string)]

							obj := ast.CreateObject(describeSObjectResultType)
							obj.Extra["info"] = sObj
							return obj
						},
					},
				},
				"newSObject": {
					&ast.Method{
						Name:       "newSObject",
						Modifiers:  []*ast.Modifier{ast.PublicModifier()},
						Parameters: []*ast.Parameter{},
						ReturnType: SObjectType,
						NativeFunction: func(this *ast.Object, parameter []*ast.Object, extra map[string]interface{}) interface{} {
							typeName := this.Extra["type"].(string)
							classType, ok := PrimitiveClassMap().Get(typeName)
							if !ok {
								panic("not found")
							}
							return ast.CreateObject(classType)
						},
					},
				},
			},
		},
	)
	classMap.Set("SObjectType", schemaSObjectType)

	describeSObjectResultType = ast.CreateClass(
		"DescribeSObjectResult",
		nil,
		&ast.MethodMap{
			Data: map[string][]*ast.Method{},
		},
		ast.NewMethodMap(),
	)
	classMap.Set("DescribeSObjectResult", describeSObjectResultType)

	sObjectTypeFields := ast.CreateClass(
		"SObjectTypeFields",
		nil,
		ast.NewMethodMap(),
		ast.NewMethodMap(),
	)
	classMap.Set("SObjectTypeFields", sObjectTypeFields)

	nameSpaceStore.Set("Schema", classMap)
}
