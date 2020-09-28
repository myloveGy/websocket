package utils

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

// Result 请求接口的返回数据
type Result struct {
	Code    int    `json:"result"`
	Content string `json:"content"`
}

// GetHTTP 发送 http 请求
func GetHTTP(message string) (*Result, error) {
	client := &http.Client{Timeout: 5 * time.Second}
	resp, err := client.Get("http://api.qingyunke.com/api.php?key=free&appid=0&msg=" + message)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	result := &Result{}
	str, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		return result, err2
	}

	err = json.Unmarshal(str, result)
	if err != nil {
		return result, err
	}

	return result, nil
}
