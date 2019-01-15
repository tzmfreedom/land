package builtin

import "github.com/tzmfreedom/goland/ast"

func init() {
	instanceMethods := NewMethodMap()

	singleEmailMessageType := CreateClass(
		"SingleEmailMessage",
		[]*Method{
			{
				Modifiers: []ast.Node{
					&ast.Modifier{
						Name: "public",
					},
				},
				Parameters: []ast.Node{},
				NativeFunction: func(this *Object, params []*Object, extra map[string]interface{}) interface{} {
					return nil
				},
			},
		},
		instanceMethods,
		nil,
	)
	singleEmailMessageType.InstanceFields.Set(
		"toAddresses",
		CreateField("toAddresses", CreateListTypeRef(stringTypeRef)),
	)
	singleEmailMessageType.InstanceFields.Set(
		"subject",
		CreateField("subject", stringTypeRef),
	)
	singleEmailMessageType.InstanceFields.Set(
		"plainTextBody",
		CreateField("plainTextBody", stringTypeRef),
	)
	singleEmailMessageType.InstanceFields.Set(
		"htmlBody",
		CreateField("htmlBody", stringTypeRef),
	)

	sendMailResultType := CreateClass(
		"SendEmailResult",
		[]*Method{
			{
				Modifiers: []ast.Node{
					&ast.Modifier{
						Name: "public",
					},
				},
				Parameters: []ast.Node{},
				NativeFunction: func(this *Object, params []*Object, extra map[string]interface{}) interface{} {
					return nil
				},
			},
		},
		NewMethodMap(),
		nil,
	)
	sendMailResultType.InstanceFields.Set(
		"success",
		CreateField("success", booleanTypeRef),
	)
	singleEmailMessageType.InstanceFields.Set(
		"errors",
		CreateField("errors", CreateListTypeRef(stringTypeRef)),
	)

	classMap := NewClassMap()
	classMap.Set("SingleEmailMessage", singleEmailMessageType)
	classMap.Set("SendEmailResult", sendMailResultType)

	nameSpaceStore.Set("Messaging", classMap)

	staticMethods := NewMethodMap()
	staticMethods.Set(
		"sendEmail",
		[]*Method{
			CreateMethod(
				"sendEmail",
				CreateListTypeRef(&ast.TypeRef{
					Name:       []string{"Messaging", "SendEmailResult"},
					Parameters: []ast.Node{},
				}),
				[]ast.Node{
					CreateListTypeParameter(&ast.TypeRef{
						// Name:       []string{"Messaging", "Email"},
						Name:       []string{"Messaging", "SingleEmailMessage"}, // TODO: implement
						Parameters: []ast.Node{},
					}),
				},
				func(this *Object, params []*Object, extra map[string]interface{}) interface{} {
					// TODO: implment
					obj := CreateObject(singleEmailMessageType)
					obj.InstanceFields.Set("errors", NewString("hoge"))
					obj.InstanceFields.Set("success", NewBoolean(true))
					listObj := CreateListObject(ListType, []*Object{obj})
					return listObj
				},
			),
		},
	)
	messagingClass := CreateClass(
		"Messaging",
		nil,
		nil,
		staticMethods,
	)
	primitiveClassMap.Set("Messaging", messagingClass)
}
