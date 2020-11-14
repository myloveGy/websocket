package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"websocket/api/response"
	"websocket/global"
	"websocket/models"
	"websocket/request"
	"websocket/utils"
)

func (a *Api) Login(c *gin.Context) {
	params := &request.UserLogin{}
	if isError, err := utils.BindAndValid(c, params); isError {
		response.NewResponseError(c, "userLogin", "参数错误:"+err.Error())
		return
	}

	// 需要查询用户是否存在
	user := &models.User{}
	if err := user.FindByUsername(global.DB, params.Username); err != nil {
		response.NewResponseError(c, "userLogin", "登录账户或者密码错误")
		return
	}

	// 验证密码是否正确
	if user.Password != params.Password {
		response.NewResponseError(c, "userLogin", "登录账户或者密码错误")
		return
	}

	// 生成redis
	accessToken := fmt.Sprintf("%d", time.Now().Unix())
	if err := a.userCache.Set(accessToken, user); err != nil {
		response.NewResponseError(c, "userLogin", "登录失败(redis error)")
		return
	}

	// 修改表
	user.AccessToken = accessToken
	if _, err := user.UpdateAccessToken(global.DB); err != nil {
		response.NewResponseError(c, "userLogin", "登录失败(db error)")
		return
	}

	response.Success(c, user)
}
