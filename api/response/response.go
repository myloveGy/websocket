package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type responseError struct {
	Code    string `json:"code"`    // 错误码
	Message string `json:"message"` // 错误描述信息
}

// 返回错误
func NewResponseError(c *gin.Context, code, message string) {
	c.JSON(http.StatusOK, responseError{Code: code, Message: message})
}

func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, data)
}
