package json

import (
	"encoding/json"

	"github.com/tidwall/gjson"
)

func JSONValid(data []byte) bool {
	return json.Valid(data)
}
func GetFieldString(data string, fieldPath string) string {
	return gjson.Get(data, fieldPath).Str
}

func GetFieldInt64(data string, fieldPath string) int64 {
	return gjson.Get(data, fieldPath).Int()
}
func GetField(data string, fieldPath string) gjson.Result {
	return gjson.Get(data, fieldPath)
}
