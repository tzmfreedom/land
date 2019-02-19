package builtin

import "github.com/tzmfreedom/land/ast"

func init() {
	instanceMethods := ast.NewMethodMap()

	singleEmailMessageType := ast.CreateClass(
		"SingleEmailMessage",
		[]*ast.Method{
			{
				Modifiers:  []*ast.Modifier{ast.PublicModifier()},
				Parameters: []*ast.Parameter{},
				NativeFunction: func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					return nil
				},
			},
		},
		instanceMethods,
		nil,
	)
	singleEmailMessageType.InstanceFields.Set(
		"toAddresses",
		ast.CreateField("toAddresses", CreateListType(StringType)),
	)
	singleEmailMessageType.InstanceFields.Set(
		"subject",
		ast.CreateField("subject", StringType),
	)
	singleEmailMessageType.InstanceFields.Set(
		"plainTextBody",
		ast.CreateField("plainTextBody", StringType),
	)
	singleEmailMessageType.InstanceFields.Set(
		"htmlBody",
		ast.CreateField("htmlBody", StringType),
	)

	sendMailResultType := ast.CreateClass(
		"SendEmailResult",
		[]*ast.Method{
			{
				Modifiers:  []*ast.Modifier{ast.PublicModifier()},
				Parameters: []*ast.Parameter{},
				NativeFunction: func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					return nil
				},
			},
		},
		ast.NewMethodMap(),
		nil,
	)
	sendMailResultType.InstanceFields.Set(
		"success",
		ast.CreateField("success", BooleanType),
	)
	singleEmailMessageType.InstanceFields.Set(
		"errors",
		ast.CreateField("errors", CreateListType(StringType)),
	)

	classMap := ast.NewClassMap()
	classMap.Set("SingleEmailMessage", singleEmailMessageType)
	classMap.Set("SendEmailResult", sendMailResultType)

	nameSpaceStore.Set("Messaging", classMap)

	staticMethods := ast.NewMethodMap()
	staticMethods.Set(
		"sendEmail",
		[]*ast.Method{
			ast.CreateMethod(
				"sendEmail",
				CreateListType(sendMailResultType),
				[]*ast.Parameter{
					CreateListTypeParameter(singleEmailMessageType),
				},
				func(this *ast.Object, params []*ast.Object, extra map[string]interface{}) interface{} {
					// TODO: implment
					obj := ast.CreateObject(singleEmailMessageType)
					obj.InstanceFields.Set("errors", NewString("hoge"))
					obj.InstanceFields.Set("success", NewBoolean(true))
					listObj := CreateListObject(ListType, []*ast.Object{obj})
					return listObj
				},
			),
		},
	)
	messagingClass := ast.CreateClass(
		"Messaging",
		nil,
		nil,
		staticMethods,
	)
	primitiveClassMap.Set("Messaging", messagingClass)
}
