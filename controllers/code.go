package controllers

type ResponseCode int64

const (
	CODESUCCESS         ResponseCode = 1000 + iota // 成功
	CODEINVALIDPARAM                               // 非法参数
	CODEUSEREXIST                                  // 用户已存在
	CODEUSERNOTEXIST                               // 用户不存在
	CODEINVALIDPASSWORD                            // 密码错误
	CODESERVERBUSY                                 // 服务器繁忙
	CODEINVALIDTOKEN                               // token无效
	CODENEEDLOGIN                                  // 需要登陆
)

var codeMessageMap = map[ResponseCode]string{
	CODESUCCESS:         "成功",
	CODEINVALIDPARAM:    "请求参数非法",
	CODEUSEREXIST:       "用户已存在",
	CODEUSERNOTEXIST:    "用户不存在",
	CODEINVALIDPASSWORD: "用户名或密码错误",
	CODESERVERBUSY:      "服务器繁忙",
	CODEINVALIDTOKEN:    "无效的token",
	CODENEEDLOGIN:       "需要登陆",
}

func (c ResponseCode) Msg() string {
	msg, ok := codeMessageMap[c]
	if !ok {
		return "未知错误"
	}
	return msg
}
