package mysql

import "errors"

var (
	ErrorUserExist          = errors.New("用户已经存在")
	ErrorUserNotExist       = errors.New("用户不存在")
	ErrorInvalidPassword    = errors.New("用户名或密码错误")
	ErrorInvalidCommunityID = errors.New("无效的社区ID")
)
