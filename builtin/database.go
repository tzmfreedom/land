package builtin

import (
	"github.com/tzmfreedom/goland/ast"
)

var saveResultType *ast.ClassType
var queryLocatorType = ast.CreateClass(
	"QueryLocator",
	[]*ast.Method{},
	ast.NewMethodMap(),
	ast.NewMethodMap(),
)

func init() {
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

	method = ast.CreateMethod(
		"getQueryLocator",
		queryLocatorType,
		[]*ast.Parameter{stringTypeParameter},
		func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
			return Null
		},
	)
	staticMethods.Set("getQueryLocator", []*ast.Method{method})

	databaseClass := ast.CreateClass(
		"Database",
		[]*ast.Method{},
		nil,
		staticMethods,
	)
	primitiveClassMap.Set("Database", databaseClass)

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
	classMap.Set("QueryLocator", queryLocatorType)

	batchableContext := ast.CreateClass(
		"BatchableContext",
		[]*ast.Method{},
		ast.NewMethodMap(),
		ast.NewMethodMap(),
	)
	classMap.Set("BatchableContext", batchableContext)

	batchableContextTypeParameter := &ast.Parameter{
		Type: batchableContext,
		Name: "_",
	}

	instanceMethods = ast.NewMethodMap()
	instanceMethods.Set(
		"start",
		[]*ast.Method{
			ast.CreateMethod(
				"start",
				queryLocatorType,
				[]*ast.Parameter{
					batchableContextTypeParameter,
				},
				nil,
			),
		},
	)
	instanceMethods.Set(
		"execute",
		[]*ast.Method{
			ast.CreateMethod(
				"execute",
				nil,
				[]*ast.Parameter{
					batchableContextTypeParameter,
					CreateListTypeParameter(T1type),
				},
				nil,
			),
		},
	)
	instanceMethods.Set(
		"finish",
		[]*ast.Method{
			ast.CreateMethod(
				"finish",
				nil,
				[]*ast.Parameter{
					batchableContextTypeParameter,
				},
				nil,
			),
		},
	)

	batchable := ast.CreateClass(
		"Batchable",
		[]*ast.Method{},
		instanceMethods,
		nil,
	)
	classMap.Set("Batchable", batchable)

	nameSpaceStore.Set("Database", classMap)
}
