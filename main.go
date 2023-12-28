package main

import (
    "bluebell/dao/mysql"
    "bluebell/dao/redis"
    "bluebell/logger"
    "bluebell/pkg/snowflake"
    "bluebell/routers"
    "bluebell/settings"
    "context"
    "fmt"
    "github.com/gin-gonic/gin"
    "go.uber.org/zap"
    "log"
    "net/http"
    "os"
    "os/signal"
    "syscall"
    "time"
)

func main() {
    // 1. init settings
    if err := settings.InitSettings(); err != nil {
        fmt.Printf("init settings failed, err: %v\n", err)
        return
    }

    // 2. init log
    if err := logger.InitLogger(settings.Config.LogConfig); err != nil {
        zap.L().Fatal("init logger failed, err: ", zap.Error(err))
        return
    }
    defer zap.L().Sync()

    // 3. init mysql
    if err := mysql.InitMySQL(settings.Config.MySQLConfig); err != nil {
        zap.L().Fatal("init mysql failed, err: ", zap.Error(err))
        return
    }
    defer mysql.Close()

    // 4. init redis
    if err := redis.InitRedis(settings.Config.RedisConfig); err != nil {
        zap.L().Fatal("init redis failed, err: ", zap.Error(err))
        return
    }
    defer redis.Close()

    // init snowflake algorithm
    if err := snowflake.Init(settings.Config.StartTime, settings.Config.MachineID); err != nil {
        zap.L().Fatal("init snowflake failed, err: ", zap.Error(err))
        return
    }

    // 5. register router
    r := routers.SetupRouter()
    r.GET("/", func(c *gin.Context) {
        c.String(http.StatusOK, settings.Config.Version)
    })

    // 6. start service
    srv := &http.Server{
        Addr:    fmt.Sprintf(":%d", settings.Config.Port),
        Handler: r,
    }

    go func() {
        if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
            log.Fatalf("listen: %s\n", err)
        }
    }()

    quit := make(chan os.Signal, 1)
    signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
    <-quit
    zap.L().Info("Shutdown Server ...")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    if err := srv.Shutdown(ctx); err != nil {
        zap.L().Fatal("Server Shutdown: ", zap.Error(err))
    }

    zap.L().Info("Server exiting")
}
