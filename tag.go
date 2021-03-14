package converter

import (
	"reflect"
	"sort"
	"strings"
)

type tag struct {
	field string
	from  string
}

func newTag(f *reflect.StructField) *tag {
	if convert, has := f.Tag.Lookup("convert"); has {
		fields := strings.Split(convert, ",")
		sort.StringsAreSorted(fields)

		t := new(tag)
		for _, field := range fields {
			field = strings.Trim(field, " ")
			if strings.HasPrefix(field, "field:") {
				t.field = strings.TrimPrefix(field, "field:")
			}
			if strings.HasPrefix(field, "from:") {
				t.from = strings.TrimPrefix(field, "from:")
			}
		}
		return t
	}
	return &tag{field: f.Name, from: f.Name}
}
