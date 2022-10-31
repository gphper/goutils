package arrutil

import (
	"errors"
	"reflect"
)

func ArrColumn(arr interface{}, fieldName string) (result []interface{}, err error) {

	defer func() {
		if msg := recover(); msg != nil {
			err = errors.New(msg.(string))
		}
	}()

	ty := reflect.TypeOf(arr)

	switch ty.Kind() {
	case reflect.Slice:
	case reflect.Array:
	default:
		return nil, errors.New("invalid type")
	}

	t := reflect.ValueOf(arr)
	if t.Len() == 0 {
		return
	}

	result = make([]interface{}, t.Len())
	for i := 0; i < t.Len(); i++ {
		result[i] = t.Index(i).FieldByName(fieldName).Interface()
	}

	return
}
