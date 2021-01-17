package utils

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnique(t *testing.T) {
	s := Unique("")
	fmt.Println(s, len(s))
	assert.Equal(t, 16, len(s))
}
