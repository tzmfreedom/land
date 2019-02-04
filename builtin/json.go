package builtin

import (
	"encoding/json"

	"fmt"

	"github.com/tzmfreedom/goland/ast"
)

func init() {
	staticMethods := ast.NewMethodMap()
	staticMethods.Set(
		"serialize",
		[]*ast.Method{
			ast.CreateMethod(
				"serialize",
				StringType,
				[]*ast.Parameter{objectTypeParameter},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					mapValue := serializeJson(params[0], false)
					value, err := json.Marshal(mapValue)
					if err != nil {
						panic(err)
					}
					return NewString(string(value))
				},
			),
			ast.CreateMethod(
				"serialize",
				StringType,
				[]*ast.Parameter{
					objectTypeParameter,
					booleanTypeParameter,
				},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					mapValue := serializeJson(params[0], params[1].BoolValue())
					value, err := json.Marshal(mapValue)
					if err != nil {
						panic(err)
					}
					return NewString(string(value))
				},
			),
		},
	)
	staticMethods.Set(
		"deserialize",
		[]*ast.Method{
			ast.CreateMethod(
				"deserialize",
				StringType,
				[]*ast.Parameter{stringTypeParameter},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					srcMap := map[string]interface{}{}
					err := json.Unmarshal([]byte(params[0].StringValue()), &srcMap)
					if err != nil {
						panic(err)
					}
					return deserializeJson(srcMap)
				},
			),
		},
	)

	classType := ast.CreateClass(
		"JSON",
		[]*ast.Method{},
		nil,
		staticMethods,
	)

	primitiveClassMap.Set("JSON", classType)
}

func serializeJson(object *ast.Object, suppressApexObjectNull bool) interface{} {
	classType := object.ClassType
	switch classType {
	case StringType:
		return object.StringValue()
	case IntegerType:
		return object.IntegerValue()
	case DoubleType:
		return object.DoubleValue()
	case BooleanType:
		return object.BoolValue()
	case NullType:
		return nil
	}
	if classType.Name == "List" {
		records := object.Extra["records"].([]*ast.Object)
		values := make([]interface{}, len(records))
		for i, record := range records {
			values[i] = serializeJson(record, suppressApexObjectNull)
		}
		return values
	}
	if classType.Name == "Map" {
		ret := map[string]interface{}{}
		values := object.Extra["values"].(map[string]*ast.Object)
		for field, value := range values {
			ret[field] = serializeJson(value, suppressApexObjectNull)
		}
		return ret
	}
	ret := map[string]interface{}{}
	for field, value := range object.InstanceFields.All() {
		if value == Null && suppressApexObjectNull {
			continue
		}
		ret[field] = serializeJson(value, suppressApexObjectNull)
	}
	return ret
}

func deserializeJson(value interface{}) *ast.Object {
	if value == nil {
		return Null
	}
	switch typedValue := value.(type) {
	case string:
		return NewString(typedValue)
	case int:
		return NewInteger(typedValue)
	case float64:
		return NewDouble(typedValue)
	case bool:
		return NewBoolean(typedValue)
	case []interface{}:
		classType := CreateListType(ObjectType)
		newObj := ast.CreateObject(classType)
		records := make([]*ast.Object, len(typedValue))
		for i, record := range typedValue {
			records[i] = deserializeJson(record)
		}
		newObj.Extra["records"] = records
		return newObj
	case map[string]interface{}:
		classType := CreateMapType(StringType, ObjectType)
		newObj := ast.CreateObject(classType)
		values := map[string]*ast.Object{}
		for key, value := range typedValue {
			values[key] = deserializeJson(value)
		}
		newObj.Extra["values"] = values
		return newObj
	}
	panic(fmt.Sprintf("no expected type %v", value))
}
