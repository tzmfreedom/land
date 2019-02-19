package builtin

import (
	"github.com/tzmfreedom/land/ast"
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

	staticMethods.Set("insert", []*ast.Method{
		ast.CreateMethod(
			"insert",
			saveResultType,
			[]*ast.Parameter{SObjectTypeParameter},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				obj := params[0]
				records := []*ast.Object{obj}
				return DatabaseDriver.Execute("insert", obj.ClassType.Name, records, "")
			},
		),
		ast.CreateMethod(
			"insert",
			saveResultType,
			[]*ast.Parameter{
				{
					Type: CreateListType(SObjectType),
					Name: "_",
				},
			},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				obj := params[0]
				records := obj.Extra["records"].([]*ast.Object)
				sObjectType := records[0].ClassType.Name
				return DatabaseDriver.Execute("insert", sObjectType, records, "")
			},
		),
	})

	staticMethods.Set("update", []*ast.Method{
		ast.CreateMethod(
			"update",
			saveResultType,
			[]*ast.Parameter{objectTypeParameter},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				obj := params[0]
				records := []*ast.Object{obj}
				return DatabaseDriver.Execute("update", obj.ClassType.Name, records, "")
			},
		),
		ast.CreateMethod(
			"update",
			saveResultType,
			[]*ast.Parameter{
				{
					Type: CreateListType(SObjectType),
					Name: "_",
				},
			},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				obj := params[0]
				records := obj.Extra["records"].([]*ast.Object)
				sObjectType := records[0].ClassType.Name
				return DatabaseDriver.Execute("update", sObjectType, records, "")
			},
		),
	})

	staticMethods.Set("delete", []*ast.Method{
		ast.CreateMethod(
			"delete",
			saveResultType,
			[]*ast.Parameter{objectTypeParameter},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				obj := params[0]
				records := []*ast.Object{obj}
				return DatabaseDriver.Execute("delete", obj.ClassType.Name, records, "")
			},
		),
		ast.CreateMethod(
			"delete",
			saveResultType,
			[]*ast.Parameter{
				{
					Type: CreateListType(SObjectType),
					Name: "_",
				},
			},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				obj := params[0]
				records := obj.Extra["records"].([]*ast.Object)
				sObjectType := records[0].ClassType.Name
				return DatabaseDriver.Execute("delete", sObjectType, records, "")
			},
		),
	})

	staticMethods.Set("upsert", []*ast.Method{
		ast.CreateMethod(
			"upsert",
			saveResultType,
			[]*ast.Parameter{
				SObjectTypeParameter,
				stringTypeParameter,
			},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				obj := params[0]
				key := params[1].StringValue()
				records := []*ast.Object{obj}
				return DatabaseDriver.Execute("upsert", obj.ClassType.Name, records, key)
			},
		),
		ast.CreateMethod(
			"upsert",
			saveResultType,
			[]*ast.Parameter{
				{
					Type: CreateListType(SObjectType),
					Name: "_",
				},
				stringTypeParameter,
			},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				obj := params[0]
				key := params[1].StringValue()
				records := obj.Extra["records"].([]*ast.Object)
				sObjectType := records[0].ClassType.Name
				return DatabaseDriver.Execute("upsert", sObjectType, records, key)
			},
		),
	})

	method := ast.CreateMethod(
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
	batchable.Interface = true
	classMap.Set("Batchable", batchable)

	nameSpaceStore.Set("Database", classMap)
}
