package dao

import "errors"

//封装每一个error,避免重复写
var (
	ErrorUserExit        = errors.New("用户已经存在")
	ErrorUserNotExit     = errors.New("用户不存在")
	ErrorInvalidPassword = errors.New("密码错误")
	ErrorInvalidID       = errors.New("无效的ID")
)
