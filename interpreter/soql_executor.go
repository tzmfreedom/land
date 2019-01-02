package interpreter

import (
	"os"

	"github.com/tzmfreedom/go-soapforce"
	"github.com/tzmfreedom/goland/ast"
	"github.com/tzmfreedom/goland/builtin"
)

type SoqlExecutor struct {
}

func (e *SoqlExecutor) Execute(n *ast.Soql) (*builtin.Object, error) {
	visitor := &ast.TosVisitor{}
	r, err := n.Accept(visitor)
	if err != nil {
		return nil, err
	}
	soql := r.(string)
	soql = soql[1 : len(soql)-1]

	result, err := executeQuery(soql)
	if err != nil {
		return nil, err
	}
	return e.getListFromResponse(n, result.Records)
}

func (e *SoqlExecutor) getListFromResponse(n *ast.Soql, records []*soapforce.SObject) (*builtin.Object, error) {
	objects := make([]*builtin.Object, len(records))
	classType, _ := builtin.PrimitiveClassMap().Get(n.FromObject)
	for i, r := range records {
		object := &builtin.Object{}
		object.ClassType = classType
		object.InstanceFields = builtin.NewObjectMap()
		object.InstanceFields.Set("id", newString(r.Id))
		for k, v := range r.Fields {
			switch val := v.(type) {
			case string:
				object.InstanceFields.Set(k, newString(val))
			}
		}
		objects[i] = object
	}
	// TODO: implement
	list := &builtin.Object{
		ClassType:      builtin.ListType,
		InstanceFields: builtin.NewObjectMap(),
		GenericType:    []*builtin.ClassType{classType},
		Extra: map[string]interface{}{
			"records": objects,
		},
	}
	return list, nil
}

func executeQuery(soql string) (*soapforce.QueryResult, error) {
	client := soapforce.NewClient()
	username := os.Getenv("SALESFORCE_USERNAME")
	password := os.Getenv("SALESFORCE_PASSWORD")
	endpoint := os.Getenv("SALESFORCE_ENDPOINT")
	client.SetLoginUrl(endpoint)
	client.Login(username, password)
	return client.Query(soql)
}