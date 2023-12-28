package mysql

import (
    "bluebell/settings"
    "fmt"
    _ "github.com/go-sql-driver/mysql"
    "github.com/jmoiron/sqlx"
    "go.uber.org/zap"
)

var db *sqlx.DB

func InitMySQL(conf *settings.MySQLConfig) (err error) {
    dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
        conf.User,
        conf.Password,
        conf.Host,
        conf.Port,
        conf.Dbname,
    )

    db, err = sqlx.Connect("mysql", dsn)
    if err != nil {
        return err
    }
    db.SetMaxIdleConns(conf.MaxIdleConns)
    db.SetMaxOpenConns(conf.MaxOpenConns)

    return err
}

func Close() {
    err := db.Close()
    if err != nil {
        zap.L().Fatal("close mysql failed: ", zap.Error(err))
    }
}
