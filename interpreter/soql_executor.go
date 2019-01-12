package interpreter

import (
	"github.com/tzmfreedom/go-soapforce"
	"github.com/tzmfreedom/goland/ast"
	"github.com/tzmfreedom/goland/builtin"
)

type SoqlExecutor struct{}

func (e *SoqlExecutor) Execute(n *ast.Soql, visitor ast.Visitor) (*builtin.Object, error) {
	records := builtin.DatabaseDriver.Query(n, visitor)
	return e.getListFromResponse(n, records)
}

func getRecords(n *ast.Soql, records []*soapforce.SObject) []*builtin.Object {
	objects := make([]*builtin.Object, len(records))
	classType, _ := builtin.PrimitiveClassMap().Get(n.FromObject)
	for i, r := range records {
		object := &builtin.Object{}
		object.ClassType = classType
		object.InstanceFields = builtin.NewObjectMap()
		object.InstanceFields.Set("id", builtin.NewString(r.Id))
		for k, v := range r.Fields {
			switch val := v.(type) {
			case string:
				object.InstanceFields.Set(k, builtin.NewString(val))
			}
		}
		objects[i] = object
	}
	return objects
}

func (e *SoqlExecutor) getListFromResponse(n *ast.Soql, records []*builtin.Object) (*builtin.Object, error) {
	// TODO: implement
	classType, _ := builtin.PrimitiveClassMap().Get(n.FromObject)
	list := &builtin.Object{
		ClassType:      builtin.ListType,
		InstanceFields: builtin.NewObjectMap(),
		GenericType:    []*builtin.ClassType{classType},
		Extra: map[string]interface{}{
			"records": records,
		},
	}
	return list, nil
}

func executeQuery(soql string) (*soapforce.QueryResult, error) {
	client := builtin.NewSoapClient()
	return client.Query(soql)
}
