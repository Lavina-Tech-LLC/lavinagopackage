package lvn

import (
	"encoding/json"
	"strings"
)

func convertKeys(j json.RawMessage) json.RawMessage {
	m := make(map[string]json.RawMessage)
	a := []json.RawMessage{}
	if err := json.Unmarshal([]byte(j), &a); err == nil {
		//JSON array object
		for k, v := range a {
			a[k] = convertKeys(v)
		}
		bytes, _ := json.Marshal(a)
		return json.RawMessage(bytes)
	}
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
	if strings.ToUpper(key) == key {
		return strings.ToLower(key)
	}
	return strings.ToLower(key[:1]) + key[1:]
}
