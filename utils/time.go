package utils

import "time"

const (
	// DateTimeLayout 默认日期时间格式
	DateTimeLayout = "2006-01-02 15:04:05"

	// DateLayout 默认日期时间格式
	DateLayout = "2006-01-02"

	// TimeZone 时区
	TimeZone = "Asia/Shanghai"
)

var loc, _ = time.LoadLocation(TimeZone)

// DateTime 当前日期时间
func DateTime() string {
	return time.Now().In(loc).Format(DateTimeLayout)
}

// Date 当前日期
func Date() string {
	return time.Now().In(loc).Format(DateLayout)
}
