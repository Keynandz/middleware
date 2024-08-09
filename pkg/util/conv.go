package util

import (
	"encoding/json"
	"strconv"
)

func ToJSON(data interface{}) string {
	json, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	return string(json)
}

func StrToUint(str string) (uint, error) {
	i, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		return 0, err
	}
	return uint(i), nil
}
