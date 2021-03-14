package converter

import (
	"reflect"
)

func Convert(from interface{}, to interface{}) {
	fromer := makeFromer(from)

	toType := reflect.TypeOf(to).Elem()
	toValue := reflect.ValueOf(to).Elem()
	for i := 0; i < toType.NumField(); i++ {
		f := toType.Field(i)
		t := newTag(&f)
		toValue.Field(i).Set(fromer[t.from])
	}
}

func makeFromer(from interface{}) map[string]reflect.Value {
	fromer := make(map[string]reflect.Value)
	fromType := reflect.TypeOf(from).Elem()
	fromValue := reflect.ValueOf(from).Elem()

	for i := 0; i < fromType.NumField(); i++ {
		f := fromType.Field(i)
		tag := newTag(&f)
		fromer[tag.field] = fromValue.Field(i)
	}

	return fromer
}
