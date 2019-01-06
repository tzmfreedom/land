package builtin

import (
	"os"

	"github.com/tzmfreedom/go-soapforce"
	"github.com/tzmfreedom/goland/ast"
)

type Sobject struct {
	Name          string
	Custom        bool
	CustomSetting bool
	Label         string
	Fields        []SobjectField
}

type SobjectField struct {
	Name             string
	Type             string
	Label            string
	RelationshipName string
	Custom           bool
	ReferenceTo      []string
}

func NewSoapClient() *soapforce.Client {
	client := soapforce.NewClient()
	username := os.Getenv("SALESFORCE_USERNAME")
	password := os.Getenv("SALESFORCE_PASSWORD")
	endpoint := os.Getenv("SALESFORCE_ENDPOINT")
	client.SetLoginUrl(endpoint)
	client.Login(username, password)
	return client
}

var typeMapper = map[string]string{
	"string":                     "String",
	"picklist":                   "String",
	"multipicklist":              "String",
	"combobox":                   "String",
	"reference":                  "String",
	"boolean":                    "Boolean",
	"currency":                   "double",
	"textarea":                   "String",
	"int":                        "Double",
	"double":                     "Double",
	"percent":                    "Double",
	"id":                         "String",
	"date":                       "Date",
	"datetime":                   "Date",
	"time":                       "Time",
	"url":                        "String",
	"email":                      "String",
	"encryptedstring":            "String",
	"datacategorygroupreference": "String",
	"location":                   "String",
	"address":                    "String",
	"anyType":                    "String",
	"complexvalue":               "String",
}

type Loader interface {
	Load() (map[string]Sobject, error)
}

func Load(src string) {
	loader := newMetaFileLoader(src)
	// TODO: sObject declaration
	sobjects, err := loader.Load()
	if err != nil {
		panic(err)
	}
	for name, sobj := range sobjects {
		fields := NewFieldMap()
		for _, f := range sobj.Fields {
			fields.Set(f.Name, &Field{
				Type: &ast.TypeRef{
					Name: []string{
						typeMapper[f.Type],
					},
				},
				Name: f.Name,
				Modifiers: []ast.Node{
					&ast.Modifier{
						Name: "public",
					},
				},
			})
		}
		primitiveClassMap.Set(name, &ClassType{
			Name:           sobj.Name,
			InstanceFields: fields,
		})
	}
}