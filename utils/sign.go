package utils

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func GetMapStringSort(data map[string]interface{}) []string {
	result := []string{}
	for k, _ := range data {
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
		case float32:
			resultString += strconv.FormatFloat(float64(data[v].(float32)), 'f', 2, 32)
		case float64:
			resultString += strconv.FormatFloat(data[v].(float64), 'f', 2, 64)
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
