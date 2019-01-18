package builtin

import (
	"github.com/tzmfreedom/go-soapforce"
	"github.com/tzmfreedom/goland/ast"
)

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
	saveResult := ast.CreateClass(
		"SaveResult",
		[]*ast.Method{},
		instanceMethods,
		ast.NewMethodMap(),
	)
	classMap.Set("SaveResult", saveResult)
	nameSpaceStore.Set("Database", classMap)

	staticMethods := ast.NewMethodMap()
	method := ast.CreateMethod(
		"insert",
		saveResult,
		[]*ast.Parameter{objectTypeParameter}, // TODO: SObject or List<SObject>
		func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
			sobj := params[0]
			record := &soapforce.SObject{}
			for k, v := range sobj.InstanceFields.All() {
				record.Fields[k] = v.Value()
			}
			client := NewSoapClient()
			rawSaveResults, err := client.Create([]*soapforce.SObject{record})
			if err != nil {
				panic(err)
			}
			retSaveResults := make([]*ast.Object, len(rawSaveResults))
			for i, sr := range rawSaveResults {
				obj := ast.CreateObject(saveResult)
				obj.Extra["isSuccess"] = NewBoolean(sr.Success)
				obj.Extra["id"] = NewString(sr.Id)
				obj.Extra["errors"] = sr.Errors
				retSaveResults[i] = obj
			}
			listObject := ast.CreateObject(ListType)
			listObject.Extra["records"] = retSaveResults
			return listObject
		},
	)
	staticMethods.Set("insert", []*ast.Method{method})
	method = ast.CreateMethod(
		"setSavePoint",
		saveResult, // TODO: implement
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
