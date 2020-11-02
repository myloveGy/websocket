package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"websocket/utils"
)

func ApiLogin(c *gin.Context) {
	info := map[string]interface{}{
		"username": "jinxing.liu",
		"user_id":  1,
		"time":     utils.DateTime(),
	}

	c.JSON(http.StatusOK, info)
}
