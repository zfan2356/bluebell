package logic

import (
	"bluebell/dao/mysql"
	"bluebell/models"
	jwt2 "bluebell/pkg/jwt"
	"bluebell/pkg/snowflake"
)

// SignUp 用户注册业务逻辑处理
func SignUp(p *models.ParamSignUp) (err error) {
	if err = mysql.CheckUserExist(p.Username); err != nil {
		// 数据库查询出错或者用户已经存在
		return err
	}
	// 生成ID
	userID := snowflake.GenID()
	u := &models.User{
		UserID:   userID,
		UserName: p.Username,
		Password: p.Password,
	}
	// 存储数据库,返回错误
	return mysql.InsertUser(u)
}

// Login 用户登陆业务逻辑处理
func Login(p *models.ParamLogin) (token string, err error) {
	var user *models.User
	user, err = mysql.FindUserByName(p.Username)
	if err != nil {
		return "", err
	}
	if user.Password != mysql.EncryptPassword(p.Password) {
		return "", mysql.ErrorInvalidPassword
	}
	// 这个时候要生成JWT的token
	return jwt2.GenToken(user.UserID, user.UserName)
}
