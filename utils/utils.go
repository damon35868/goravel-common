package utils

import "encoding/json"

func MarshalStr(data any) string {
	jsonData, _ := json.Marshal(data)
	return string(jsonData)
}
