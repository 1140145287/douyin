package controller

type ResCode int64

const CtxUserIdKey = "user_id"

const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExit
	CodeUserNotExit
	CodeInvalidPassword
	CodeServerBusy
	CodeEmptyAuth
	CodePresentationError
	CodeInvalidAuth
	CodeInternalError
)

var CodeMap = map[ResCode]string{
	CodeSuccess:           "success",
	CodeInvalidParam:      "请求参数错误",
	CodeUserExit:          "用户已存在",
	CodeUserNotExit:       "用户不存在",
	CodeInvalidPassword:   "用户名或密码错误",
	CodeServerBusy:        "服务繁忙",
	CodeEmptyAuth:         "请求头中auth为空",
	CodePresentationError: "请求头中auth格式有误",
	CodeInvalidAuth:       "无效的Token",
	CodeInternalError:     "内部服务错误",
}

func (c ResCode) Msg() string {
	msg, ok := CodeMap[c]
	if !ok {
		msg = CodeMap[CodeServerBusy]
	}
	return msg
}
