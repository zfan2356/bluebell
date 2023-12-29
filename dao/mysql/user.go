package mysql

import (
	"bluebell/models"
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
)

const secret string = "zfan2356"

var (
	ErrorUserExist    = errors.New("用户已经存在")
	ErrorUserNotExist = errors.New("用户不存在")
)

// InsertUser 插入用户记录,
func InsertUser(user *models.User) (err error) {
	// 对密码进行加密
	sqlStr := "insert into user(user_id, username, password) values (?, ?, ?)"
	_, err = db.Exec(sqlStr, user.UserID, user.UserName, EncryptPassword(user.Password))
	return err
}

// CheckUserExist 检查指定用户名的用户是否存在
func CheckUserExist(name string) error {
	sqlStr := "select count(user_id) from user where username = ?"

	var count int
	if err := db.Get(&count, sqlStr, name); err != nil {
		return err
	}
	if count > 0 {
		return ErrorUserExist
	}
	return nil
}

// EncryptPassword 加密算法, 利用hash值进行加密
func EncryptPassword(orgPassword string) string {
	hash := md5.New()
	hash.Write([]byte(secret))
	return hex.EncodeToString(hash.Sum([]byte(orgPassword)))
}

func FindUserByName(username string) (user *models.User, err error) {
	sqlStr := "select user_id, password, username from user where username = ?"
	user = new(models.User)
	err = db.Get(user, sqlStr, username)
	if errors.Is(err, sql.ErrNoRows) {
		return user, ErrorUserNotExist
	}
	return
}
