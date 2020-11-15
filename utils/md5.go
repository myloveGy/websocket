package utils

import (
	"crypto/md5"
	"fmt"
)

func Md5(s string) string {
	has := md5.Sum([]byte(s))
	return fmt.Sprintf("%x", has)
}
