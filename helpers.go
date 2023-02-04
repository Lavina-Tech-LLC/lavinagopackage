package lvn

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"reflect"
	"syscall"

	"github.com/iancoleman/orderedmap"
)

func isNull[T any](n T) bool {
	switch any(n).(type) {
	case bool:
		return !any(n).(bool)
	case int:
		return any(n).(int) == 0
	case int32:
		return any(n).(int32) == 0
	case int64:
		return any(n).(int64) == 0
	case float32:
		return any(n).(float32) == 0
	case float64:
		return any(n).(float64) == 0
	case string:
		return any(n).(string) == ""
	default:
		return any(n) == nil || reflect.ValueOf(n).IsNil()
	}
}

func Ternary[T any, N any](condition N, ifTrue, ifFalse T) T {
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

func WaitExitSignal() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	fmt.Println("Shuting down Server ...")
}
