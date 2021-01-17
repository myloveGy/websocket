package user

import (
	"github.com/gin-gonic/gin"
	"websocket/api/response"
	"websocket/request/admin"
	"websocket/utils"
)

func (u *User) Create(c *gin.Context) {
	// 解析参数
	params := &admin.UserCreate{}
	if isError, err := utils.BindAndValid(c, params); isError {
		response.InvalidParams(c, err)
		return
	}

	// 查询数据
	if _, err := u.userService.Create(params); err != nil {
		response.BusinessError(c, err)
		return
	}

	response.Success(c, params)
	return
}
