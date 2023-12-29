package routers

import (
	"bluebell/controllers"
	"bluebell/logger"
	"bluebell/pkg/translator"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func SetupRouter(mode string) *gin.Engine {
	if mode == gin.ReleaseMode {
		gin.SetMode(gin.ReleaseMode)
	}
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))

	// 将校验器注册为中文
	if err := translator.InitTrans("zh"); err != nil {
		zap.L().Error("init translator failed", zap.Error(err))
	}
	//注册业务路由
	r.POST("/signup", controllers.SignUpHandler)
	r.POST("/login", controllers.LoginHandler)

	return r
}
