package lvn

import (
	"strings"

	"github.com/iancoleman/orderedmap"
	"golang.org/x/exp/slices"
)

func convertKeys(o interface{}, omitKeys, selectKeys []string) interface{} {

	switch t := o.(type) {
	case orderedmap.OrderedMap:
		newMap := orderedmap.New()
		keys := t.Keys()
		for _, k := range keys {
			fixed := fixKey(k)
			v, _ := t.Get(k)
			// checking selectKeys
			if len(selectKeys) != 0 && !slices.Contains(selectKeys, fixed) {
				continue
			}

			// checking omitKeys
			if len(omitKeys) != 0 && slices.Contains(omitKeys, fixed) {
				continue
			}
			newMap.Set(fixed, convertKeys(v, omitKeys, selectKeys))
		}
		return *newMap
	case []interface{}:
		newArray := []interface{}{}
		for _, v := range t {
			newArray = append(newArray, convertKeys(v, omitKeys, selectKeys))
		}
		return newArray
	default:
		return t
	}
}

func fixKey(key string) string {
	if strings.ToUpper(key) == key {
		return strings.ToLower(key)
	}
	return strings.ToLower(key[:1]) + key[1:]
}
