package builtin

import "github.com/tzmfreedom/goland/ast"

var messageType = ast.CreateClass(
	"Message",
	[]*ast.Method{
		{
			Modifiers: []*ast.Modifier{ast.PublicModifier()},
			Parameters: []*ast.Parameter{
				severityTypeParameter,
				stringTypeParameter,
			},
			NativeFunction: func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				this.InstanceFields.Set("severity", params[0])
				this.InstanceFields.Set("summary", params[1])
				return nil
			},
		},
		{
			Modifiers: []*ast.Modifier{ast.PublicModifier()},
			Parameters: []*ast.Parameter{
				severityTypeParameter,
				stringTypeParameter,
				stringTypeParameter,
			},
			NativeFunction: func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				this.InstanceFields.Set("severity", params[0])
				this.InstanceFields.Set("summary", params[1])
				this.InstanceFields.Set("detail", params[2])
				return nil
			},
		},
		{
			Modifiers: []*ast.Modifier{ast.PublicModifier()},
			Parameters: []*ast.Parameter{
				severityTypeParameter,
				stringTypeParameter,
				stringTypeParameter,
				stringTypeParameter,
			},
			NativeFunction: func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				this.InstanceFields.Set("severity", params[0])
				this.InstanceFields.Set("summary", params[1])
				this.InstanceFields.Set("detail", params[2])
				this.InstanceFields.Set("id", params[3])
				return nil
			},
		},
	},
	ast.NewMethodMap(),
	ast.NewMethodMap(),
)

var messageTypeParameter = &ast.Parameter{
	Name: "_",
	Type: messageType,
}

var severityType = ast.CreateEnum("Severity", []string{"CONFIRM", "ERROR", "FATAL", "INFO", "WARNING"})

var severityTypeParameter = &ast.Parameter{
	Name: "_",
	Type: severityType,
}

func init() {
	instanceMethods := ast.NewMethodMap()
	staticMethods := ast.NewMethodMap()
	staticMethods.Set(
		"currentPage",
		[]*ast.Method{
			ast.CreateMethod(
				"currentPage",
				pageReferenceType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					return extra["current_page"]
				},
			),
		},
	)
	staticMethods.Set(
		"addMessage",
		[]*ast.Method{
			ast.CreateMethod(
				"addMessage",
				nil,
				[]*ast.Parameter{
					messageTypeParameter,
				},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					if extra["messages"] == nil {
						extra["messages"] = CreateListObject(CreateListType(messageType), []*ast.Object{})
					}
					messages := extra["messages"].(*ast.Object)
					records := messages.Extra["records"].([]*ast.Object)
					messages.Extra["records"] = append(records, params[0])
					return nil
				},
			),
		},
	)
	staticMethods.Set(
		"addMessages",
		[]*ast.Method{
			ast.CreateMethod(
				"addMessages",
				nil,
				[]*ast.Parameter{
					exceptionTypeParameter,
				},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					messages := extra["messages"].(*ast.Object)
					records := messages.Extra["records"].([]*ast.Object)
					messageObject := ast.CreateObject(messageType)
					messageObject.Extra["value"] = params[0].Extra["message"]
					messages.Extra["records"] = append(records, messageObject)
					return nil
				},
			),
		},
	)
	staticMethods.Set(
		"getMessages",
		[]*ast.Method{
			ast.CreateMethod(
				"getMessages",
				CreateListType(messageType),
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					return extra["messages"]
				},
			),
		},
	)
	staticMethods.Set(
		"hasMessages",
		[]*ast.Method{
			ast.CreateMethod(
				"hasMessages",
				pageReferenceType,
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					if extra["messages"] == nil {
						return NewBoolean(false)
					}
					messages := extra["messages"].(*ast.Object)
					records := messages.Extra["records"].([]*ast.Object)
					return NewInteger(len(records))
				},
			),
		},
	)

	classType := ast.CreateClass(
		"ApexPages",
		[]*ast.Method{},
		instanceMethods,
		staticMethods,
	)

	primitiveClassMap.Set("ApexPages", classType)

	messageType.InstanceFields.Set("severity", ast.CreateField("severity", severityType))
	messageType.InstanceFields.Set("summary", ast.CreateField("summary", StringType))
	messageType.InstanceFields.Set("detail", ast.CreateField("detail", StringType))
	messageType.InstanceFields.Set("id", ast.CreateField("id", StringType))
	messageType.InstanceMethods.Set("getDetail", []*ast.Method{
		ast.CreateMethod(
			"getDetail",
			StringType,
			[]*ast.Parameter{},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				value, ok := this.InstanceFields.Get("getDetail")
				if !ok {
					panic("no instance field")
				}
				return value
			},
		),
	})
	messageType.InstanceMethods.Set("getSeverity", []*ast.Method{
		ast.CreateMethod(
			"getSeverity",
			severityType,
			[]*ast.Parameter{},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				value, ok := this.InstanceFields.Get("getSeverity")
				if !ok {
					panic("no instance field")
				}
				return value
			},
		),
	})
	messageType.InstanceMethods.Set("getSummary", []*ast.Method{
		ast.CreateMethod(
			"getSummary",
			StringType,
			[]*ast.Parameter{},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				value, ok := this.InstanceFields.Get("getSummary")
				if !ok {
					panic("no instance field")
				}
				return value
			},
		),
	})
	messageType.InstanceMethods.Set("getComponentLabel", []*ast.Method{
		ast.CreateMethod(
			"getComponentLabel",
			StringType,
			[]*ast.Parameter{},
			func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
				value, ok := this.InstanceFields.Get("getComponentLabel")
				if !ok {
					panic("no instance field")
				}
				return value
			},
		),
	})

	account, _ := PrimitiveClassMap().Get("Account")
	instanceMethods = ast.NewMethodMap()
	instanceMethods.Set(
		"getRecord",
		[]*ast.Method{
			ast.CreateMethod(
				"getRecord",
				account, // TODO: SObject
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					return this.Extra["record"]
				},
			),
		},
	)
	instanceMethods.Set(
		"getId",
		[]*ast.Method{
			ast.CreateMethod(
				"getId",
				StringType, // TODO: SObject
				[]*ast.Parameter{},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					record := this.Extra["record"].(*ast.Object)
					value, _ := record.InstanceFields.Get("Id")
					return value
				},
			),
		},
	)
	accountType, _ := PrimitiveClassMap().Get("Account")

	staticMethods = ast.NewMethodMap()
	classType = ast.CreateClass(
		"StandardController",
		[]*ast.Method{
			{
				Modifiers: []*ast.Modifier{ast.PublicModifier()},
				Parameters: []*ast.Parameter{
					{
						Type: accountType, // TODO: sobject
						Name: "_",
					},
				},
				NativeFunction: func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					this.Extra["record"] = params[0]
					return nil
				},
			},
		},
		instanceMethods,
		staticMethods,
	)

	classMap := ast.NewClassMap()
	classMap.Set("StandardController", classType)
	classMap.Set("Message", messageType)
	classMap.Set("Severity", severityType)

	nameSpaceStore.Set("ApexPages", classMap)
}
