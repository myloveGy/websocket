package user

import (
	"github.com/gin-gonic/gin"
	"websocket/api/response"
	"websocket/config"
	"websocket/request/admin"
	"websocket/service/api"
	"websocket/utils"
)

type User struct {
	userService *api.UserService
}

func NewUser(userService *api.UserService) *User {
	return &User{userService: userService}
}

func (u *User) List(c *gin.Context) {
	// 解析参数
	params := &admin.UserSearch{}
	if isError, err := utils.BindAndValid(c, params); isError {
		response.InvalidParams(c, err)
		return
	}

	if params.Page <= 0 {
		params.Page = 1
	}

	if params.PageSize == 0 {
		params.PageSize = config.App.DefaultPageSize
	}

	// 查询数据
	users, total, err := u.userService.List(params)
	if err != nil {
		response.BusinessError(c, err)
	}

	response.PageSuccess(c, users, total, params.Page)
	return
}
