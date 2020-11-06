package middleware

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type SignParam struct {
	AppId int    `form:"app_id" json:"app_id" binding:"required"`
	Time  string `form:"time" json:"time" binding:"required"`
	Sign  string `form:"sign" json:"sign" binding:"required"`
}

func Sign() gin.HandlerFunc {
	return func(context *gin.Context) {
		data, _ := ioutil.ReadAll(context.Request.Body)
		mapData := map[string]interface{}{}
		json.Unmarshal(data, &mapData)
		if emptyKey, isEmpty := verifyEmptyKeys(mapData, []string{"app_id", "time", "sign"}); isEmpty {
			fmt.Println("123", emptyKey+" is Empty")
			context.Abort()
		}

		context.Next()
	}
}

func verifyEmptyKeys(data map[string]interface{}, keys []string) (string, bool) {
	for _, v := range keys {
		if _, ok := data[v]; !ok {
			return v, true
		}
	}

	return "", false
}
