package redis

import (
    "bluebell/settings"
    "fmt"
    "github.com/go-redis/redis"
    "go.uber.org/zap"
)

var (
    rdb *redis.Client
)

func InitRedis(conf *settings.RedisConfig) (err error) {
    rdb = redis.NewClient(&redis.Options{
        Addr: fmt.Sprintf("%s:%d",
            conf.Host,
            conf.Port),
        Password: conf.Password,
        PoolSize: conf.PoolSize,
        DB:       conf.DB,
    })

    _, err = rdb.Ping().Result()
    return err
}

func Close() {
    err := rdb.Close()
    if err != nil {
        zap.L().Fatal("close redis failed: ", zap.Error(err))
    }
}
