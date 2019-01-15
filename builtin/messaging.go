package builtin

import (
	"github.com/tzmfreedom/goland/ast"
)

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
		CreateField("toAddresses", stringTypeRef),
	)
	singleEmailMessageType.InstanceFields.Set(
		"plainTextBody",
		CreateField("toAddresses", stringTypeRef),
	)
	singleEmailMessageType.InstanceFields.Set(
		"htmlBody",
		CreateField("toAddresses", stringTypeRef),
	)

	sendMailResultType := CreateClass(
		"SendMailResult",
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
	classMap.Set("SendMailResult", sendMailResultType)

	nameSpaceStore.Set("Messaging", classMap)
}
