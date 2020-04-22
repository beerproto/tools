package mapping

import (
	"encoding/json"
	"reflect"
	"strings"
)

func Enum(m map[string]json.RawMessage) (map[string]json.RawMessage) {



	return m
}

func IterateFields(i interface{}) []string {
	fields := make([]string, 0)
	t := reflect.TypeOf(i)
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		j, ok := field.Tag.Lookup("json")
		if !ok {
			continue
		}

		fragments := strings.Split(j, ",")
		if len(fragments) > 0 {
			fields = append(fields, fragments[0])
		}
	}
	return fields
}