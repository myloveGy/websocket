package utils

import (
	"strings"
	"websocket/models"

	"github.com/gin-gonic/gin"
	ut "github.com/go-playground/universal-translator"
	val "github.com/go-playground/validator/v10"
)

type ValidError struct {
	Key     string
	Message string
}

type ValidErrors []*ValidError

func (v *ValidError) Error() string {
	return v.Message
}

func (v ValidErrors) Error() string {
	return strings.Join(v.Errors(), ",")
}

func (v ValidErrors) Errors() []string {
	var errs []string
	for _, err := range v {
		errs = append(errs, err.Error())
	}

	return errs
}

func BindAndValid(c *gin.Context, v interface{}) (bool, ValidErrors) {
	var errs ValidErrors
	err := c.ShouldBind(v)
	if err != nil {
		v := c.Value("trans")
		trans, _ := v.(ut.Translator)
		validErrors, ok := err.(val.ValidationErrors)
		if !ok {
			return true, nil
		}

		for key, value := range validErrors.Translate(trans) {
			errs = append(errs, &ValidError{
				Key:     key,
				Message: value,
			})
		}

		return true, errs
	}

	return false, nil
}

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
