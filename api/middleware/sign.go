package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"websocket/api/response"
	"websocket/entity"
	"websocket/utils"
)

type SignParam struct {
	AppId int    `form:"app_id" json:"app_id" binding:"required"`
	Time  string `form:"time" json:"time" binding:"required"`
	Sign  string `form:"sign" json:"sign" binding:"required"`
}

func (m *MiddleWare) Sign() gin.HandlerFunc {
	return func(context *gin.Context) {
		body, _ := ioutil.ReadAll(context.Request.Body)
		data := map[string]interface{}{}
		context.Request.Body = ioutil.NopCloser(bytes.NewReader(body))
		// 需要传递json数据
		if err := json.Unmarshal(body, &data); err != nil {
			response.InvalidParams(context, "请求格式错误(必须json)")
			return
		}

		// 验证必须传递 app_id time, sign
		if emptyKey, isEmpty := utils.VerifyEmptyKeys(data, []string{"app_id", "time", "sign"}); isEmpty {
			response.InvalidParams(context, "参数:"+emptyKey+" is empty")
			return
		}

		// 拿到应用id
		appId, ok := data["app_id"].(string)
		if !ok {
			response.InvalidParams(context, "参数: app_id 必须为字符串")
			return
		}

		// 获取到应用信息
		app, err := m.appRepo.FindByAppId(appId)
		if err != nil {
			response.Unauthorized(context, "应用信息错误")
		}

		// 验证签名信息
		if utils.Sign(data, app.AppSecret) != data["sign"].(string) {
			response.Unauthorized(context, "签名信息错误")
			return
		}

		// 验证应用状态
		if err != nil || app.Status != entity.AppStatusActivate {
			response.Unauthorized(context, "应用信息错误")
			return
		}

		context.Set("app", app)
		context.Next()
	}
}
