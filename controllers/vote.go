package controllers

import (
	"bluebell/logic"
	"bluebell/models"
	"bluebell/pkg/translator"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// PostVoteHandler 投票响应函数
func PostVoteHandler(c *gin.Context) {
	p := new(models.ParamVoteData)
	if err := c.ShouldBindJSON(p); err != nil {
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			zap.L().Error("vote invalid param", zap.Error(err))
			ResponseError(c, CODEINVALIDPARAM)
			return
		}
		errt := translator.RemoveTopStruct(errs.Translate(translator.Trans))
		ResponseErrorWithMsg(c, CODEINVALIDPARAM, errt)
		return
	}

	userID, err := GetCurrentUser(c)
	if err != nil {
		ResponseError(c, CODENEEDLOGIN)
		return
	}

	if err := logic.VoteForPost(userID, p); err != nil {
		zap.L().Error("logic.VoteForPost(userID, p) failed",
			zap.Int64("userID", userID),
			zap.Int64("post_id", p.PostID),
			zap.Error(err),
		)
		ResponseError(c, CODESERVERBUSY)
		return
	}
	ResponseSuccess(c, nil)
}
