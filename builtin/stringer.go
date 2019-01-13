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
	if o.ClassType.ToString != nil {
		return o.ClassType.ToString(o)
	}
	switch o.ClassType {
	case ListType:
		records := o.Extra["records"].([]*Object)
		recordExpressions := make([]string, len(records))
		v.AddIndent(func() {
			for i, record := range records {
				recordExpressions[i] = v.withIndent(v.String(record))
			}
		})
		recordsString := ""
		if len(recordExpressions) != 0 {
			recordsString = "\n" + strings.Join(recordExpressions, ",\n") + "\n"
		}
		return fmt.Sprintf(
			`<List> {%s%s`,
			recordsString,
			v.withIndent("}"),
		)
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
	if o.ClassType == nil {
		return "null"
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
