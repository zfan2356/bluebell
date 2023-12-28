package routers

import (
    "bluebell/logger"
    "github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
    r := gin.New()
    r.Use(logger.GinLogger(), logger.GinRecovery(true))
    return r
}
