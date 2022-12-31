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

// 删除指定元素
func Unset[T int | string](items []T, item T) []T {
	result := make([]T, len(items)-1)

	m := 0
	for i := 0; i < len(items); i++ {
		if items[i] == item {
			continue
		}
		result[m] = items[i]
		m++
	}

	return result
}
