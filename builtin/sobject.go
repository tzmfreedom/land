package builtin

import (
	"fmt"
	"sort"
	"strings"

	"github.com/tzmfreedom/go-soapforce"
	"github.com/tzmfreedom/land/ast"
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
			ToString:        SObjectType.ToString,
		})
	}
}

var SObjectType = &ast.ClassType{Name: "SObject"}
var SObjectTypeParameter = &ast.Parameter{
	Type: SObjectType,
	Name: "_",
}

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

	SObjectType.Constructors = []*ast.Method{}
	SObjectType.InstanceFields = ast.NewFieldMap()
	SObjectType.StaticFields = ast.NewFieldMap()
	SObjectType.InstanceMethods = instanceMethods
	SObjectType.StaticMethods = ast.NewMethodMap()
	SObjectType.ToString = func(o *ast.Object) string {
		i := 0
		names := make([]string, len(o.InstanceFields.Data))
		for name, _ := range o.InstanceFields.Data {
			names[i] = name
			i++
		}
		sort.Slice(names, func(i, j int) bool {
			return names[i] < names[j]
		})

		fields := []string{}
		for _, name := range names {
			r, ok := o.InstanceFields.Get(name)
			if !ok {
				panic("InstanceFields.Get#failed")
			}
			if r == Null {
				continue
			}
			fields = append(fields, fmt.Sprintf("  %s: %s", name, String(r)))
		}
		fieldStr := ""
		if len(fields) != 0 {
			fieldStr = fmt.Sprintf("\n%s\n", strings.Join(fields, "\n"))
		}

		return fmt.Sprintf(
			`<%s> {%s%s`,
			o.ClassType.Name,
			fieldStr,
			"}",
		)
	}
	primitiveClassMap.Set("SObject", SObjectType)
}
