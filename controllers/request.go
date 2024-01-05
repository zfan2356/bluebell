package controllers

import (
	"bluebell/models"
	"errors"
	"github.com/gin-gonic/gin"
	"strconv"
)

var (
	ErrorUserNotLogin = errors.New("用户未登陆")
)

// GetCurrentUser 获取当前登陆的userID
func GetCurrentUser(c *gin.Context) (userID int64, err error) {
	uid, ok := c.Get(models.ContextUserIDKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userID, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
	}
	return
}

// GetPageAndSizeInfo 返回查询时需要的分页值和页内大小限制, 非法时默认第一页, 大小为10
func GetPageAndSizeInfo(c *gin.Context) (int64, int64) {
	var (
		page int64
		size int64
		err  error
	)
	page, err = strconv.ParseInt(c.Query("page"), 10, 64)
	if err != nil {
		page = 1
	}
	size, err = strconv.ParseInt(c.Query("size"), 10, 64)
	if err != nil {
		size = 10
	}
	return page, size
}
