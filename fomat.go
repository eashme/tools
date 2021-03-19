package tools

import (
	jsoniter "github.com/json-iterator/go"
	"reflect"
	"strconv"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

func JsonMarshal(v interface{}) ([]byte, error) {
	return json.Marshal(v)
}

func JsonUnmarshal(b []byte, v interface{}) error {
	return json.Unmarshal(b, v)
}

func Json(v interface{}) string {
	b, _ := json.Marshal(v)
	return string(b)
}

func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

func Str2Int64(s string) int64 {
	i, _ := strconv.ParseInt(s, 10, 64)
	return i
}
