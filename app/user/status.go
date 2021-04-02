package user

import (
	"sync"
)

// err文件定义用户模块中一些错误信息，登录和注册错误
// 错误码从以10000开始
// 正确码从20000开始

type StatusCode int64

var (
	mutex = &sync.Mutex{}
)

// 用户模块响应请求的状态的错误码
const (
	ErrUserNotExist StatusCode = 10000 + iota
	ErrUserExist
	ErrInvalidUsernameOrPassword
	ErrInvalidParam
)

// 用户模块相关的成功响应请求的状态码
const (
	LonginSuccess StatusCode = 20000 + iota
	SignupSuccess
)

// userStatus 保存所有用户状态信息
var userStatusMap = map[StatusCode]string{
	// 错误码
	ErrUserNotExist:              "用户不存在",
	ErrUserExist:                 "用户已经存在",
	ErrInvalidUsernameOrPassword: "用户名或者密码错误",

	// 正确码
	LonginSuccess: "用户登录成功",
	SignupSuccess: "用户注册成功",
}

// GetStatusString 获取状态对应的字符串
func GetStatusString(code StatusCode) string {
	mutex.Lock()
	defer mutex.Unlock()
	return userStatusMap[code]
}
