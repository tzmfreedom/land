package builtin

import (
	"fmt"
	"sort"
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
		names := make([]string, len(o.InstanceFields.Data))
		for name, _ := range o.InstanceFields.Data {
			names[i] = name
			i++
		}
		sort.Slice(names, func(i, j int) bool {
			return names[i] < names[j]
		})

		for i, name := range names {
			obj := ""
			v.AddIndent(func() {
				r, _ := o.InstanceFields.Get(name)
				obj = v.String(r)
			})
			fields[i] = v.withIndent(fmt.Sprintf("%s: %s", name, obj))
		}
	})
	fieldStr := ""
	if len(fields) != 0 {
		fieldStr = fmt.Sprintf("\n%s\n", strings.Join(fields, "\n"))
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
