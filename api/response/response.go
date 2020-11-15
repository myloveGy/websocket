package response

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"websocket/entity"
)

type errorItem struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func NewError(c *gin.Context, code int, message ...interface{}) {
	httpStatus, ok := entity.CodeHttpStatusMap[code]
	if !ok {
		httpStatus = http.StatusInternalServerError
	}

	// 判断错误类型
	var msg string
	if message != nil {
		switch message[0].(type) {
		case error:
			msg = message[0].(error).Error()
		case string:
			msg = message[0].(string)
		default:
			msg = fmt.Sprintf("%v", message[0])
		}
	}

	// 没有消息内容
	if msg == "" {
		if msg1, ok := entity.CodeMessageMap[code]; ok {
			msg = msg1
		} else {
			msg = entity.CodeMessageMap[entity.CodeSystemError]
		}
	}

	c.JSON(httpStatus, errorItem{
		Code:    code,
		Message: msg,
	})
	c.Abort()
}

func InvalidParams(c *gin.Context, message ...interface{}) {
	NewError(c, entity.CodeInvalidParamsError, message...)
}

func Unauthorized(c *gin.Context, message ...interface{}) {
	NewError(c, entity.CodeUnauthorizedError, message...)
}

func NotLogin(c *gin.Context, message ...interface{}) {
	NewError(c, entity.CodeNotLoginError, message...)
}

func BusinessError(c *gin.Context, message ...interface{}) {
	NewError(c, entity.CodeBusinessError, message...)
}

func SystemError(c *gin.Context, message ...interface{}) {
	NewError(c, entity.CodeSystemError, message...)
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}
