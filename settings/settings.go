package settings

import (
    "fmt"
    "github.com/fsnotify/fsnotify"
    "github.com/spf13/viper"
)

var (
    Config = new(AppConfig)
)

type AppConfig struct {
    Name         string `mapstructure:"name"`
    Mode         string `mapstructure:"mode"`
    Port         int    `mapstructure:"port"`
    Version      string `mapstructure:"version"`
    StartTime    string `mapstructure:"start_time"`
    MachineID    int64  `mapstructure:"machine_id"`
    *MySQLConfig `mapstructure:"mysql"`
    *LogConfig   `mapstructure:"log"`
    *RedisConfig `mapstructure:"redis"`
}

type MySQLConfig struct {
    Host         string `mapstructure:"host"`
    Port         int    `mapstructure:"port"`
    User         string `mapstructure:"user"`
    Password     string `mapstructure:"password"`
    Dbname       string `mapstructure:"dbname"`
    MaxOpenConns int    `mapstructure:"max_open_conns"`
    MaxIdleConns int    `mapstructure:"max_idle_conns"`
}
type LogConfig struct {
    Level      string `mapstructure:"level"`
    FileName   string `mapstructure:"filename"`
    MaxSize    int    `mapstructure:"max_size"`
    MaxAge     int    `mapstructure:"max_age"`
    MaxBackups int    `mapstructure:"max_backups"`
}
type RedisConfig struct {
    Host     string `mapstructure:"host"`
    Port     int    `mapstructure:"port"`
    DB       int    `mapstructure:"db"`
    Password string `mapstructure:"password"`
    PoolSize int    `mapstructure:"pool_size"`
}

func InitSettings() (err error) {
    viper.SetConfigFile("config.yaml")
    viper.AddConfigPath("./")

    if err = viper.ReadInConfig(); err != nil {
        fmt.Printf("viper readconfig failed, err = %v\n", err)
        return err
    }

    if err = viper.Unmarshal(Config); err != nil {
        fmt.Printf("viper unmarshal failed, err = %v\n", err)
        return err
    }

    viper.WatchConfig()
    viper.OnConfigChange(func(in fsnotify.Event) {
        fmt.Printf("viper config changed\n")
        if err = viper.Unmarshal(Config); err != nil {
            fmt.Printf("viper unmarshal failed, err = %v\n", err)
        }
    })
    return err
}
