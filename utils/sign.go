package utils

import (
	"crypto/md5"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

func GetMapStringSort(data map[string]interface{}) []string {
	result := []string{}
	for k := range data {
		result = append(result, k)
	}

	sort.Sort(sort.StringSlice(result))
	return result
}

func MapToString(data map[string]interface{}) string {
	result := GetMapStringSort(data)
	var resultString string
	for _, v := range result {
		if v == "sign" {
			continue
		}

		resultString += v + "="
		switch data[v].(type) {
		case bool:
			resultString += strconv.FormatBool(data[v].(bool))
		case int:
			resultString += strconv.Itoa(data[v].(int))
		case int64:
			resultString += strconv.FormatInt(data[v].(int64), 10)
		case string:
			resultString += data[v].(string)
		default:
			resultString += fmt.Sprint(data[v])
		}

		resultString += "&"
	}

	return strings.Trim(resultString, "&")
}

func Sign(data map[string]interface{}, Secret string) string {
	linkString := MapToString(data)
	w := md5.New()
	_, _ = io.WriteString(w, linkString+Secret)
	return fmt.Sprintf("%x", w.Sum(nil))
}

func VerifyEmptyKeys(data map[string]interface{}, keys []string) (string, bool) {
	for _, v := range keys {
		if _, ok := data[v]; !ok {
			return v, true
		}
	}

	return "", false
}
