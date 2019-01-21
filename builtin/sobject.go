package builtin

import (
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

func NewSoapClient(username, password, endpoint string) *soapforce.Client {
	if soapClient != nil {
		return soapClient
	}
	soapClient := soapforce.NewClient()
	soapClient.SetLoginUrl(endpoint)
	_, err := soapClient.Login(username, password)
	if err != nil {
		panic(err)
	}
	return soapClient
}

var typeMapper = map[string]*ast.ClassType{
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

func LoadSObjectClass(src string) {
	loader := newMetaFileLoader(src)
	// TODO: sObject declaration
	var err error
	sObjects, err = loader.Load()
	if err != nil {
		panic(err)
	}
	for name, sobj := range sObjects {
		fields := ast.NewFieldMap()
		for _, f := range sobj.Fields {
			fields.Set(f.Name, &ast.Field{
				Type:      typeMapper[f.Type],
				Name:      f.Name,
				Modifiers: []*ast.Modifier{ast.PublicModifier()},
			})
		}
		primitiveClassMap.Set(name, &ast.ClassType{
			Name:           sobj.Name,
			Constructors:   []*ast.Method{},
			InstanceFields: fields,
		})
	}
}
