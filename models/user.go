package models

// User 用户
type User struct {
	UserID   int64  `db:"user_id"`
	UserName string `db:"username"`
	Password string `db:"password"`
	Token    string
}
