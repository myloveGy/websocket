package user

import (
	"github.com/gin-gonic/gin"
	"strconv"
	"websocket/api/response"
	"websocket/helpers"
	"websocket/utils"
)

type User struct {
}

func NewUser() *User {
	return &User{}
}

func (u *User) Ws(c *gin.Context) {
	user := helpers.GetUser(c)
	if user == nil {
		response.NotLogin(c)
		return
	}

	app := helpers.GetApp(c)
	if app == nil {
		response.NotLogin(c)
		return
	}

	// 需要进行加密
	data := map[string]interface{}{
		"time":    utils.DateTime(),
		"user_id": strconv.FormatInt(user.UserId, 10),
		"app_id":  app.AppId,
	}

	data["sign"] = utils.Sign(data, app.AppSecret)

	response.Success(c, data)
}
