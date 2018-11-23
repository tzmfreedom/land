package builtin

import (
	"fmt"
	"strings"
)

type ToStringer struct {
	Indent int
}

func (v *ToStringer) AddIndent(f func()) {
	v.Indent += 2
	f()
	v.Indent -= 2
}

func (v *ToStringer) withIndent(src string) string {
	return strings.Repeat(" ", v.Indent) + src
}

func (v *ToStringer) String(o *Object) string {
	switch o.ClassType {
	case IntegerType:
		return fmt.Sprintf("%d", o.Value().(int))
	case StringType:
		return o.Value().(string)
	case BooleanType:
		return fmt.Sprintf("%t", o.Value().(bool))
	case DoubleType:
		return fmt.Sprintf("%f", o.Value().(float64))
	}
	fields := make([]string, len(o.InstanceFields.All()))
	v.AddIndent(func() {
		i := 0
		for name, field := range o.InstanceFields.All() {
			obj := ""
			v.AddIndent(func() {
				obj = v.String(field)
			})
			fields[i] = v.withIndent(fmt.Sprintf("%s: %s", name, obj))
			i++
		}
	})
	fieldStr := ""
	if len(fields) != 0 {
		fieldStr = fmt.Sprintf(" \n%s\n", strings.Join(fields, "\n"))
	}
	return fmt.Sprintf(
		`<%s> {%s%s`,
		o.ClassType.Name,
		fieldStr,
		v.withIndent("}"),
	)
}

func String(o *Object) string {
	stringer := &ToStringer{}
	return stringer.String(o)
}
