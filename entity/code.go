package entity

import "net/http"

const (
	CodeInvalidParamsError = 40001 // 参数错误
	CodeUnauthorizedError  = 40002 // 未授权
	CodeNotLoginError      = 40003 // 未登录
	CodeBusinessError      = 40004 // 业务错误
	CodeSystemError        = 40005 // 服务器异常
)

var (
	// 错误码对应消息
	CodeMessageMap = map[int]string{
		CodeInvalidParamsError: "参数错误",
		CodeUnauthorizedError:  "未授权",
		CodeNotLoginError:      "未登录",
		CodeBusinessError:      "业务错误",
		CodeSystemError:        "系统错误，请重试",
	}

	// 错误码对应http状态
	CodeHttpStatusMap = map[int]int{
		CodeInvalidParamsError: http.StatusBadRequest,
		CodeUnauthorizedError:  http.StatusUnauthorized,
		CodeNotLoginError:      http.StatusPaymentRequired,
		CodeBusinessError:      http.StatusBadRequest,
	}
)
