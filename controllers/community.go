package controllers

import (
	"bluebell/logic"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"strconv"
)

// 社区相关

// CommunityHandler 社区响应函数
func CommunityHandler(c *gin.Context) {
	// 查询到所有的社区, 以列表的形式返回
	data, err := logic.GetCommunityList()
	if err != nil {
		zap.L().Error("logic.GetCommunityList() failed", zap.Error(err))
		ResponseError(c, CODESERVERBUSY)
	}
	ResponseSuccess(c, data)
}

// CommunityDetailHandler 社区分类详情响应函数
func CommunityDetailHandler(c *gin.Context) {
	// 获取社区id
	communityID, err := strconv.ParseInt(c.Param("id"), 10, 64)
	if err != nil {
		ResponseError(c, CODEINVALIDPARAM)
		return
	}
	communityDetail, err := logic.GetCommunityDetail(communityID)
	if err != nil {
		zap.L().Error("logic.GetCommunityDetail() failed", zap.Error(err))
		ResponseError(c, CODESERVERBUSY)
	}
	ResponseSuccess(c, communityDetail)
}
