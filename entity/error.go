package entity

const (
	ErrAppNoTExists       = "应用信息不存在"
	ErrApp                = "应用信息错误"
	ErrAppDisable         = "应用已经停用"
	ErrUserNotEmpty       = "用户字段信息不能为空"
	ErrCreateMessage      = "创建消息信息失败"
	ErrCreateUserMessage  = "创建用户消息失败"
	ErrDecodeReplyMessage = "解析回复消息失败"
	ErrUpdateUserMessage  = "标记用户消息已读失败"
	ErrUserNotExists      = "用户信息不存在"
	ErrUserUpdateStatus   = "修改用户状态失败"
	ErrUserCreate         = "创建用户信息失败"
	ErrUserUpdate         = "修改用户信息失败"
	ErrUserPasswordLength = "密码不能小于6位字符"
	ErrUserPasswordEncode = "密码加密失败"
	ErrUserUsernameExists = "用户名称已经存在"
	ErrUserPhoneExists    = "用户手机号已经存在"
	ErrAppCreate          = "创建应用失败"
	ErrAppUpdate          = "修改应用失败"
	ErrAppUpdateStatus    = "修改应用状态失败"
	ErrAppDelete          = "删除应用失败"
)
