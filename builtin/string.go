package builtin

import (
	"strings"

	"regexp"

	"github.com/tzmfreedom/goland/ast"
)

var stringTypeRef = &ast.TypeRef{
	Name:       []string{"String"},
	Parameters: []*ast.TypeRef{},
}

var StringType = &ast.ClassType{Name: "String"}

func createStringType(c *ast.ClassType) *ast.ClassType {
	instanceMethods := ast.NewMethodMap()
	instanceMethods.Set("replace", []*ast.Method{
		ast.CreateMethod(
			"replace",
			StringType,
			[]*ast.Parameter{
				stringTypeParameter,
				stringTypeParameter,
			},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				target := params[0].StringValue()
				replacement := params[1].StringValue()
				src := this.StringValue()
				return NewString(strings.Replace(src, target, replacement, -1))
			},
		),
	})
	instanceMethods.Set("replaceAll", []*ast.Method{
		ast.CreateMethod(
			"replaceAll",
			StringType,
			[]*ast.Parameter{
				stringTypeParameter,
				stringTypeParameter,
			},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				regExp := params[0].StringValue()
				replacement := params[1].StringValue()
				src := this.StringValue()
				r := regexp.MustCompile(regExp)
				return NewString(r.ReplaceAllString(src, replacement))
			},
		),
	})
	instanceMethods.Set("indexOf", []*ast.Method{
		ast.CreateMethod(
			"indexOf",
			StringType,
			[]*ast.Parameter{
				stringTypeParameter,
			},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				substring := params[0].StringValue()
				src := this.StringValue()
				return NewInteger(strings.Index(src, substring))
			},
		),
	})
	instanceMethods.Set("isBlank", []*ast.Method{
		ast.CreateMethod(
			"isBlank",
			BooleanType,
			[]*ast.Parameter{},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				if this == Null {
					return NewBoolean(true)
				}
				src := this.StringValue()
				return NewBoolean(strings.TrimSpace(src) == "")
			},
		),
	})
	instanceMethods.Set("isNotBlank", []*ast.Method{
		ast.CreateMethod(
			"isNotBlank",
			BooleanType,
			[]*ast.Parameter{},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				if this == Null {
					return NewBoolean(false)
				}
				src := this.StringValue()
				return NewBoolean(strings.TrimSpace(src) != "")
			},
		),
	})
	instanceMethods.Set("isEmpty", []*ast.Method{
		ast.CreateMethod(
			"isEmpty",
			BooleanType,
			[]*ast.Parameter{},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				if this == Null {
					return NewBoolean(true)
				}
				src := this.StringValue()
				return NewBoolean(src == "")
			},
		),
	})
	instanceMethods.Set("isNotEmpty", []*ast.Method{
		ast.CreateMethod(
			"isNotEmpty",
			BooleanType,
			[]*ast.Parameter{},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				if this == Null {
					return NewBoolean(false)
				}
				src := this.StringValue()
				return NewBoolean(src != "")
			},
		),
	})
	instanceMethods.Set("equals", []*ast.Method{
		ast.CreateMethod(
			"equals",
			BooleanType,
			[]*ast.Parameter{
				stringTypeParameter,
			},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				secondString := params[0].StringValue()
				src := this.StringValue()
				return NewBoolean(src == secondString)
			},
		),
	})
	instanceMethods.Set("contains", []*ast.Method{
		ast.CreateMethod(
			"contains",
			BooleanType,
			[]*ast.Parameter{
				stringTypeParameter,
			},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				substring := params[0].StringValue()
				src := this.StringValue()
				return NewBoolean(strings.Contains(src, substring))
			},
		),
	})
	instanceMethods.Set("length", []*ast.Method{
		ast.CreateMethod(
			"length",
			IntegerType,
			[]*ast.Parameter{},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				return NewInteger(len(this.StringValue()))
			},
		),
	})
	instanceMethods.Set("split", []*ast.Method{
		ast.CreateMethod(
			"split",
			CreateListType(StringType),
			[]*ast.Parameter{stringTypeParameter},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				split := params[0].StringValue()
				src := this.StringValue()
				parts := strings.Split(src, split)
				records := make([]*ast.Object, len(parts))
				for i, part := range parts {
					records[i] = NewString(part)
				}
				listType := ast.CreateObject(ListType)
				listType.Extra["records"] = records
				return listType
			},
		),
	})
	instanceMethods.Set("substring", []*ast.Method{
		ast.CreateMethod(
			"substring",
			StringType,
			[]*ast.Parameter{IntegerTypeParameter},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				startIndex := params[0].IntegerValue()
				return NewString(this.StringValue()[startIndex:])
			},
		),
		ast.CreateMethod(
			"substring",
			StringType,
			[]*ast.Parameter{
				IntegerTypeParameter,
				IntegerTypeParameter,
			},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				startIndex := params[0].IntegerValue()
				endIndex := params[1].IntegerValue()
				return NewString(this.StringValue()[startIndex:endIndex])
			},
		),
	})
	instanceMethods.Set("toLowerCase", []*ast.Method{
		ast.CreateMethod(
			"toLowerCase",
			StringType,
			[]*ast.Parameter{},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				return NewString(strings.ToLower(this.StringValue()))
			},
		),
	})
	instanceMethods.Set("toUpperCase", []*ast.Method{
		ast.CreateMethod(
			"toUpperCase",
			StringType,
			[]*ast.Parameter{},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				return NewString(strings.ToUpper(this.StringValue()))
			},
		),
	})
	staticMethods := ast.NewMethodMap()
	staticMethods.Set("join", []*ast.Method{
		ast.CreateMethod(
			"join",
			StringType,
			[]*ast.Parameter{
				CreateListTypeParameter(ObjectType),
				stringTypeParameter,
			},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				iterableObj := params[0].Extra["records"].([]*ast.Object)
				separator := params[1].StringValue()
				elements := make([]string, len(iterableObj))
				for i, obj := range iterableObj {
					elements[i] = String(obj)
				}
				return NewString(strings.Join(elements, separator))
			},
		),
	})
	staticMethods.Set("valueOf", []*ast.Method{
		ast.CreateMethod(
			"valueOf",
			StringType,
			[]*ast.Parameter{BlobTypeParameter},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				src := params[0].Extra["value"].([]byte)
				return NewString(string(src))
			},
		),
		ast.CreateMethod(
			"valueOf",
			StringType,
			[]*ast.Parameter{objectTypeParameter},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				toConvert := params[0]
				return NewString(String(toConvert))
			},
		),
	})
	c.InstanceMethods = instanceMethods
	c.StaticMethods = staticMethods
	c.ToString = func(o *ast.Object) string {
		return o.Value().(string)
	}
	return c
}

var stringTypeParameter = &ast.Parameter{
	Type: StringType,
	Name: "_",
}

func init() {
	createStringType(StringType)
	primitiveClassMap.Set("String", StringType)
}
