package api

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
	"websocket/api/response"
	"websocket/entity"
	"websocket/models"
	"websocket/request"
	"websocket/utils"
)

func (a *Api) Register(c *gin.Context) {
	params := &request.ApiRegisterParams{}
	if isError, err := utils.BindAndValid(c, params); isError {
		response.InvalidParams(c, err)
		return
	}

	// 需要查询用户是否存在
	user1, err := a.userRepo.FindByUsername(params.Username)
	if user1 != nil || !errors.Is(err, sql.ErrNoRows) {
		response.InvalidParams(c, "登录账号已经存在")
		return
	}

	// 查询手机号是否已经存在
	user2, err2 := a.userRepo.FindByPhone(params.Phone)
	if user2 != nil || !errors.Is(err2, sql.ErrNoRows) {
		response.InvalidParams(c, "手机账号已经存在")
		return
	}

	user := &models.User{
		Username:    params.Username,
		Phone:       params.Phone,
		Password:    params.Password,
		Status:      entity.AppStatusActivate,
		AccessToken: utils.Md5(fmt.Sprintf("%s:%d", params.Username, time.Now().Unix())),
	}

	// 新增入库
	if err := a.userRepo.Create(user); err != nil {
		response.SystemError(c, "新增用户数据失败")
		return
	}

	// 写入缓存
	if err := a.userCache.Set(user.AccessToken, user); err != nil {
		response.SystemError(c, "登录失败(redis error)")
		return
	}

	response.Success(c, user)
}
