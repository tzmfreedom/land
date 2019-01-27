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
	loader := NewMetaFileLoader(src)
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
			Name:            sobj.Name,
			SuperClass:      SObjectType,
			Constructors:    []*ast.Method{},
			InstanceFields:  fields,
			StaticFields:    ast.NewFieldMap(),
			InstanceMethods: ast.NewMethodMap(),
			StaticMethods:   ast.NewMethodMap(),
		})
	}
}

var SObjectType = &ast.ClassType{Name: "SObject"}

func init() {
	instanceMethods := ast.NewMethodMap()
	instanceMethods.Set(
		"put",
		[]*ast.Method{
			ast.CreateMethod(
				"put",
				nil,
				[]*ast.Parameter{
					stringTypeParameter,
					objectTypeParameter,
				},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					key := params[0].StringValue()
					this.InstanceFields.Set(key, params[1])
					return nil
				},
			),
		},
	)
	instanceMethods.Set(
		"get",
		[]*ast.Method{
			ast.CreateMethod(
				"get",
				ObjectType,
				[]*ast.Parameter{stringTypeParameter},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					key := params[0].StringValue()
					value, ok := this.InstanceFields.Get(key)
					if !ok {
						// TODO: impl
					}
					return value
				},
			),
		},
	)

	SObjectType.Constructors = []*ast.Method{
		{
			Modifiers:  []*ast.Modifier{ast.PublicModifier()},
			Parameters: []*ast.Parameter{},
			NativeFunction: func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				return nil
			},
		},
	}
	SObjectType.InstanceFields = ast.NewFieldMap()
	SObjectType.StaticFields = ast.NewFieldMap()
	SObjectType.InstanceMethods = instanceMethods
	SObjectType.StaticMethods = ast.NewMethodMap()
	primitiveClassMap.Set("SObject", SObjectType)
}
