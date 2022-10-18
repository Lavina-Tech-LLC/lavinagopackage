package lvn

import (
	"encoding/json"
	"reflect"
)

type Nullable interface {
	bool | int | string
}

func isNull[T Nullable](n T) bool {
	switch any(n).(type) {
	case bool:
		return !any(n).(bool)
	case int:
		return any(n).(int) == 0
	case string:
		return any(n).(string) == ""
	default:
		return false
	}
}

func Ternary[T any, N Nullable](condition N, ifTrue, ifFalse T) T {
	if !isNull(condition) {
		return ifTrue
	} else {
		return ifFalse
	}
}

// Marshals in camelCase
func Marshal(data any) ([]byte, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return bytes, err
	}

	return convertKeys(json.RawMessage(bytes)), nil
}

func GetValue[T any](object any, fieldNames ...string) T {
	obj := object
	for _, fn := range fieldNames {
		r := reflect.ValueOf(obj)
		val := reflect.Indirect(r).FieldByName(fn)
		obj = val.Interface()
	}
	return obj.(T)
}
