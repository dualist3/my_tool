package utils

import "encoding/json"

func ToJson(v any) string {
	jsonData, err := json.Marshal(v)
	if err != nil {
		return ""
	}
	return string(jsonData)
}
