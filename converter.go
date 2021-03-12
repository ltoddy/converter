package converter

import (
	"reflect"
)

func Convert(from interface{}, to interface{}) error {
	fromer := make(map[string]reflect.Value)
	fromType := reflect.TypeOf(from).Elem()
	fromValue := reflect.ValueOf(from).Elem()
	for i := 0; i < fromType.NumField(); i++ {
		f := fromType.Field(i)
		fromer[f.Name] = fromValue.Field(i)
	}

	toType := reflect.TypeOf(to).Elem()
	toValue := reflect.ValueOf(to).Elem()
	for i := 0; i < toType.NumField(); i++ {
		f := toType.Field(i)
		if convert, has := f.Tag.Lookup("convert"); has {
			mapping, err := newMappingFromString(convert)
			if err != nil {
				return err
			}
			toValue.Field(i).Set(fromer[mapping.from])
		}
	}

	return nil
}
