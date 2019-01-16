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

var soapClient *soapforce.Client

func NewSoapClient() *soapforce.Client {
	if soapClient != nil {
		return soapClient
	}
	soapClient := soapforce.NewClient()
	username := os.Getenv("SALESFORCE_USERNAME")
	password := os.Getenv("SALESFORCE_PASSWORD")
	endpoint := os.Getenv("SALESFORCE_ENDPOINT")
	soapClient.SetLoginUrl(endpoint)
	soapClient.Login(username, password)
	return soapClient
}

var typeMapper = map[string]*ClassType{
	"string":        StringType,
	"picklist":      StringType,
	"multipicklist": StringType,
	"combobox":      StringType,
	"reference":     StringType,
	"boolean":       BooleanType,
	"currency":      DoubleType,
	"textarea":      StringType,
	"int":           DoubleType,
	"double":        DoubleType,
	"percent":       DoubleType,
	"id":            StringType,
	//"date":                       dateType,
	//"datetime":                   dateType,
	//"time":                       TimeType,
	"url":                        StringType,
	"email":                      StringType,
	"encryptedstring":            StringType,
	"datacategorygroupreference": StringType,
	"location":                   StringType,
	"address":                    StringType,
	"anyType":                    StringType,
	"complexvalue":               StringType,
}

type Loader interface {
	Load() (map[string]Sobject, error)
}

var sObjects map[string]Sobject

func Load(src string) {
	loader := newMetaFileLoader(src)
	// TODO: sObject declaration
	var err error
	sObjects, err = loader.Load()
	if err != nil {
		panic(err)
	}
	for name, sobj := range sObjects {
		fields := NewFieldMap()
		for _, f := range sobj.Fields {
			fields.Set(f.Name, &Field{
				Type: typeMapper[f.Type],
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
			Constructors:   []*Method{},
			InstanceFields: fields,
		})
	}
}
