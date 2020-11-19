package helpers

import (
	"github.com/gin-gonic/gin"

	"websocket/models"
)

func GetUser(c *gin.Context) *models.User {
	value, ok := c.Get("user")
	if !ok {
		return nil
	}

	user, ok1 := value.(*models.User)
	if !ok1 {
		return nil
	}

	return user
}

func GetApp(c *gin.Context) *models.App {
	value, ok := c.Get("app")
	if !ok {
		return nil
	}

	app, ok1 := value.(*models.App)
	if !ok1 {
		return nil
	}

	return app
}
