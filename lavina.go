package lavina

import (
	"encoding/json"
	"strings"
)

type response struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	IsOk    bool        `json:"isOk"`
}

func Response(data interface{}, message string, status bool) response {
	result := response{}
	result.Message = message
	result.Data = data
	result.IsOk = status
	return result
}

func ResponseCamelCase(statusCode int, data interface{}, message string) (int, string, []byte) {
	result := response{}
	result.Message = message
	result.Data = data
	result.IsOk = statusCode < 300 && statusCode >= 200
	bytes, _ := json.Marshal(result)

	return statusCode, "application/json", convertKeys(json.RawMessage(bytes))
}

func convertKeys(j json.RawMessage) json.RawMessage {
	m := make(map[string]json.RawMessage)
	if err := json.Unmarshal([]byte(j), &m); err != nil {
		// Not a JSON object
		return j
	}

	for k, v := range m {
		fixed := fixKey(k)
		delete(m, k)
		m[fixed] = convertKeys(v)
	}

	b, err := json.Marshal(m)
	if err != nil {
		return j
	}

	return json.RawMessage(b)
}

func fixKey(key string) string {
	return strings.ToUpper(key)
}
