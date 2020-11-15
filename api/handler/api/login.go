package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"websocket/api/response"
	"websocket/request"
	"websocket/utils"
)

func (a *Api) Login(c *gin.Context) {
	params := &request.UserLogin{}
	if isError, err := utils.BindAndValid(c, params); isError {
		response.InvalidParams(c, err)
		return
	}

	// 需要查询用户是否存在
	user, err := a.userRepo.FindByUsername(params.Username)
	if err != nil {
		response.BusinessError(c, "登录账户或者密码错误")
		return
	}

	// 验证密码是否正确
	if user.Password != params.Password {
		response.BusinessError(c, "登录账户或者密码错误")
		return
	}

	// 删除之前的redis(目的，值允许一台设备登录)
	a.userCache.Delete(user.AccessToken)

	// 生成redis
	user.AccessToken = utils.Md5(fmt.Sprintf("%d:%d", user.UserId, time.Now().Unix()))
	if err := a.userCache.Set(user.AccessToken, user); err != nil {
		response.SystemError(c, "登录失败(redis error)")
		return
	}

	// 修改表
	if _, err := a.userRepo.UpdateAccessToken(user.UserId, user.AccessToken); err != nil {
		response.SystemError(c, "登录失败(db error)")
		return
	}

	response.Success(c, user)
}
