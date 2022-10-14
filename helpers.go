package lvn

import "encoding/json"

func Ternary(condition bool, ifTrue, ifFalse any) any {
	if condition {
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
