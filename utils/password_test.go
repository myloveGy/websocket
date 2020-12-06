package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGeneratePassword(t *testing.T) {
	password, err := GeneratePassword("123456")
	assert.NoError(t, err)
	assert.NotEmpty(t, password)
	fmt.Printf("%s\n", password)
}

func TestValidatePassword(t *testing.T) {
	isOk := ValidatePassword("123456", "$2a$10$b0rneS6JSP57AToZ5.lMe.YnKb9oIlq5/wMJIIa5X2Jk9gLXXuw7C")
	assert.Equal(t, true, isOk)

	isOk1 := ValidatePassword("1234567", "$2a$10$b0rneS6JSP57AToZ5.lMe.YnKb9oIlq5/wMJIIa5X2Jk9gLXXuw7C")
	assert.Equal(t, false, isOk1)
}
