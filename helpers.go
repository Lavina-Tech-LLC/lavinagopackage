package lvn

import (
	"encoding/json"
	"reflect"

	"github.com/iancoleman/orderedmap"
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

func marshal(data any, omitKeys, selectKeys []string) ([]byte, error) {
	bytes, err := json.Marshal(data)
	if err != nil {
		return bytes, err
	}

	o := orderedmap.New()

	err = json.Unmarshal(bytes, &o)
	if err != nil {
		return bytes, err
	}

	newO := convertKeys(*o, omitKeys, selectKeys).(orderedmap.OrderedMap)
	return json.Marshal(newO)
}

// Marshals in camelCase
func Marshal(data any, omitKeys ...string) ([]byte, error) {
	return marshal(data, omitKeys, []string{})
}

func MarshalSelected(data any, selectKeys ...string) ([]byte, error) {
	return marshal(data, []string{}, selectKeys)
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
