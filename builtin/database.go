package builtin

import (
	"github.com/tzmfreedom/goland/ast"
)

var saveResultType *ast.ClassType

func init() {
	instanceMethods := ast.NewMethodMap()
	instanceMethods.Set(
		"getErrors",
		[]*ast.Method{
			ast.CreateMethod(
				"getErrors",
				CreateListType(StringType),
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					return this.Extra["errors"]
				},
			),
		},
	)
	instanceMethods.Set(
		"getId",
		[]*ast.Method{
			ast.CreateMethod(
				"getId",
				StringType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					return this.Extra["id"]
				},
			),
		},
	)
	instanceMethods.Set(
		"isSuccess",
		[]*ast.Method{
			ast.CreateMethod(
				"isSuccess",
				BooleanType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					return this.Extra["isSuccess"]
				},
			),
		},
	)

	classMap := ast.NewClassMap()
	saveResultType = ast.CreateClass(
		"SaveResult",
		[]*ast.Method{},
		instanceMethods,
		ast.NewMethodMap(),
	)
	classMap.Set("SaveResult", saveResultType)
	nameSpaceStore.Set("Database", classMap)

	staticMethods := ast.NewMethodMap()
	method := ast.CreateMethod(
		"insert",
		saveResultType,
		[]*ast.Parameter{objectTypeParameter}, // TODO: SObject or List<SObject>
		func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
			obj := params[0]
			var records []*ast.Object
			// TODO: check SObject class
			if obj.ClassType.Name == "List" {
				records = obj.Extra["records"].([]*ast.Object)
			} else {
				records = []*ast.Object{obj}
			}
			sObjectType := records[0].ClassType.Name
			return DatabaseDriver.Execute("insert", sObjectType, records, "")
		},
	)
	staticMethods.Set("insert", []*ast.Method{method})
	method = ast.CreateMethod(
		"setSavePoint",
		nil, // TODO: implement
		[]*ast.Parameter{},
		func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
			return Null
		},
	)
	staticMethods.Set("setSavePoint", []*ast.Method{method})
	method = ast.CreateMethod(
		"rollback",
		nil,
		[]*ast.Parameter{objectTypeParameter}, // TODO: savepoint
		func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
			return Null
		},
	)
	staticMethods.Set("rollback", []*ast.Method{method})

	databaseClass := ast.CreateClass(
		"Database",
		[]*ast.Method{},
		nil,
		staticMethods,
	)
	primitiveClassMap.Set("Database", databaseClass)

}
