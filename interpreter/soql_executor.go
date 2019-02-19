package interpreter

import (
	"github.com/tzmfreedom/go-soapforce"
	"github.com/tzmfreedom/land/ast"
	"github.com/tzmfreedom/land/builtin"
)

type SoqlExecutor struct{}

func (e *SoqlExecutor) Execute(n *ast.Soql, visitor ast.Visitor) (*ast.Object, error) {
	records := builtin.DatabaseDriver.Query(n, visitor)
	return e.getListFromResponse(n, records)
}

func getRecords(n *ast.Soql, records []*soapforce.SObject) []*ast.Object {
	objects := make([]*ast.Object, len(records))
	classType, ok := builtin.PrimitiveClassMap().Get(n.FromObject)
	if !ok {
		panic(n.FromObject + "not found")
	}
	for i, r := range records {
		object := &ast.Object{}
		object.ClassType = classType
		object.InstanceFields = ast.NewObjectMap()
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

func (e *SoqlExecutor) getListFromResponse(n *ast.Soql, records []*ast.Object) (*ast.Object, error) {
	// TODO: implement
	classType, ok := builtin.PrimitiveClassMap().Get(n.FromObject)
	if !ok {
		panic(n.FromObject + "not found")
	}
	list := &ast.Object{
		ClassType:      builtin.CreateListType(classType),
		InstanceFields: ast.NewObjectMap(),
		Extra: map[string]interface{}{
			"records": records,
		},
	}
	return list, nil
}
