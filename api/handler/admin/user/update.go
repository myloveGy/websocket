package user

import (
	"github.com/gin-gonic/gin"
	"websocket/api/response"
	"websocket/request/admin"
	"websocket/utils"
)

func (u *User) Update(c *gin.Context) {
	// 解析参数
	params := &admin.UserUpdate{}
	if isError, err := utils.BindAndValid(c, params); isError {
		response.InvalidParams(c, err)
		return
	}

	// 查询数据
	if err := u.userService.Update(params); err != nil {
		response.BusinessError(c, err)
		return
	}

	response.Success(c, params)
	return
}
