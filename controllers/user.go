package controllers

import (
	"bluebell/dao/mysql"
	"bluebell/logic"
	"bluebell/models"
	"bluebell/pkg/translator"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.uber.org/zap"
)

// SignUpHandler 注册响应函数
func SignUpHandler(c *gin.Context) {
	// 处理登陆模块的功能
	// 1. 参数校验, 查看用户登陆信息是否合法
	p := new(models.ParamSignUp)
	if err := c.ShouldBindJSON(p); err != nil {
		// 请求参数有误, 直接返回响应
		// 打印日志
		zap.L().Error("SignUp with invalid params", zap.Error(err))
		// 翻译错误信息
		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CODEINVALIDPARAM)
			return
		}
		// 翻译错误
		ResponseErrorWithMsg(c, CODEINVALIDPARAM, translator.RemoveTopStruct(errs.Translate(translator.Trans)))
		return
	}

	// 2. 登陆成功, 业务处理
	if err := logic.SignUp(p); err != nil {
		if errors.Is(err, mysql.ErrorUserExist) {
			ResponseError(c, CODEUSEREXIST)
		}
		ResponseError(c, CODESERVERBUSY)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, nil)
}

// LoginHandler 登陆响应函数
func LoginHandler(c *gin.Context) {
	// 1. 获取请求参数和参数校验工作
	p := &models.ParamLogin{}
	if err := c.ShouldBindJSON(p); err != nil {
		zap.L().Error("login with invalid params", zap.Error(err))

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResponseError(c, CODEINVALIDPARAM)
			return
		}
		// 翻译错误
		ResponseErrorWithMsg(c, CODEINVALIDPARAM, translator.RemoveTopStruct(errs.Translate(translator.Trans)))
		return
	}
	// 2. 业务逻辑处理
	if err := logic.Login(p); err != nil {
		zap.L().Error("logic.Login failed", zap.Error(err))
		if errors.Is(err, mysql.ErrorUserNotExist) {
			ResponseError(c, CODEUSERNOTEXIST)
		}
		ResponseError(c, CODEINVALIDPASSWORD)
		return
	}
	// 3. 返回响应
	ResponseSuccess(c, nil)
}
