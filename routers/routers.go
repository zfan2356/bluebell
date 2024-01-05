package routers

import (
	"bluebell/controllers"
	"bluebell/logger"
	"bluebell/middleware"
	"bluebell/pkg/translator"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
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

	// v1
	v1 := r.Group("/api/v1")
	// 注册业务路由
	v1.POST("/signup", controllers.SignUpHandler)
	// 登陆业务路由
	v1.POST("/login", controllers.LoginHandler)
	// 应用JWT认证中间件
	v1.Use(middleware.JWTAuthMiddleware())
	{
		v1.GET("/community", controllers.CommunityHandler)
		v1.GET("/community/:id", controllers.CommunityDetailHandler)

		// 帖子创建
		v1.POST("/post", controllers.CreatePostHandler)
		v1.GET("/post/:id", controllers.GetPostDetailHandler)
		// 单纯获取所有的帖子, 没有其他功能
		v1.GET("/posts1", controllers.GetPostList1Handler)
		// 根据时间或者分数获取帖子列表
		v1.GET("/posts2", controllers.GetPostListWithOrderHandler)
		// 根据时间, 分数排序获取帖子列表, 并且还可以做到社区分类, 如果没有指定社区就默认全体社区
		v1.GET("/posts", controllers.GetPostListHandler)

		// 投票
		v1.POST("/vote", controllers.PostVoteHandler)
	}

	r.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "404",
		})
	})
	return r
}
