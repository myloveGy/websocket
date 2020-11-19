package utils

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestDate(t *testing.T) {
	assert.Equal(t, time.Now().Format(DateLayout), Date())
}

func TestDateTime(t *testing.T) {
	assert.Equal(t, time.Now().Format(DateTimeLayout), DateTime())
}
