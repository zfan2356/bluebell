package controllers

import (
	"bluebell/logic"
	"bluebell/models"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// CreatePostHandler 创建帖子的响应函数
func CreatePostHandler(c *gin.Context) {
	// 1. 获取参数以及参数校验
	p := new(models.Post)
	if err := c.ShouldBindJSON(p); err != nil {
		ResponseError(c, CODEINVALIDPARAM)
		return
	}
	// 2. 完善p结构体的信息, 也就是填写一些可以获取的信息
	userID, err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c, CODENEEDLOGIN)
		return
	}
	p.AuthorID = userID

	// 3. 创建帖子并存储
	if err := logic.CreatePost(p); err != nil {
		zap.L().Error("logic.CreatePost(p) failed", zap.Error(err))
		ResponseError(c, CODESERVERBUSY)
		return
	}

	// 4. 返回响应
	ResponseSuccess(c, nil)
}

// GetPostDetailHandler 获取帖子的信息
func GetPostDetailHandler(c *gin.Context) {
	// 1. 获取参数(从URL中获取帖子的ID)
	pid, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		zap.L().Error("get post detail with invalid param", zap.Error(err))
		ResponseError(c, CODEINVALIDPARAM)
		return
	}
	// 2. 根据id取出数据
	data, err := logic.GetPostById(pid)
	if err != nil {
		zap.L().Error("logic.GetPostById(pid) failed", zap.Error(err))
		ResponseError(c, CODESERVERBUSY)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, data)
}

// GetPostListHandler 获取所有帖子的响应函数
func GetPostListHandler(c *gin.Context) {
	// 获取分页和限制参数
	page, size := GetPageAndSizeInfo(c)
	data, err := logic.GetPostList(page, size)
	if err != nil {
		zap.L().Error("logic.GetPostList() failed", zap.Error(err))
		ResponseError(c, CODESERVERBUSY)
		return
	}
	ResponseSuccess(c, data)
}
