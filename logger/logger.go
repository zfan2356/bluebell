package logger

import (
    "bluebell/settings"
    "github.com/gin-gonic/gin"
    "github.com/natefinch/lumberjack"
    "go.uber.org/zap"
    "go.uber.org/zap/zapcore"
    "net"
    "net/http"
    "net/http/httputil"
    "os"
    "runtime/debug"
    "strings"
    "time"
)

func InitLogger(conf *settings.LogConfig) (err error) {
    writeSyncer := getLogWriter(
        conf.FileName,
        conf.MaxSize,
        conf.MaxBackups,
        conf.MaxAge,
    )
    encoder := getEncoder()

    var l = new(zapcore.Level)
    if err = l.UnmarshalText([]byte(conf.Level)); err != nil {
        return err
    }
    core := zapcore.NewCore(encoder, writeSyncer, l)

    // replace logger
    lg := zap.New(core, zap.AddCaller())
    zap.ReplaceGlobals(lg)

    return err
}

func getEncoder() zapcore.Encoder {
    encoderConfig := zap.NewProductionEncoderConfig()
    encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
    encoderConfig.TimeKey = "time"
    encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
    encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
    encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
    return zapcore.NewJSONEncoder(encoderConfig)
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer {
    lumberJackLogger := &lumberjack.Logger{
        Filename:   filename,
        MaxSize:    maxSize,
        MaxBackups: maxBackup,
        MaxAge:     maxAge,
    }
    return zapcore.AddSync(lumberJackLogger)
}

// GinLogger receive gin default logger
func GinLogger() gin.HandlerFunc {
    return func(c *gin.Context) {
        start := time.Now()
        path := c.Request.URL.Path
        query := c.Request.URL.RawQuery
        c.Next()

        cost := time.Since(start)
        zap.L().Info(path,
            zap.Int("status", c.Writer.Status()),
            zap.String("method", c.Request.Method),
            zap.String("path", path),
            zap.String("query", query),
            zap.String("ip", c.ClientIP()),
            zap.String("user-agent", c.Request.UserAgent()),
            zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
            zap.Duration("cost", cost),
        )
    }
}

// GinRecovery recover panic，then use zap to logger
func GinRecovery(stack bool) gin.HandlerFunc {
    return func(c *gin.Context) {
        defer func() {
            if err := recover(); err != nil {
                // Check for a broken connection, as it is not really a
                // condition that warrants a panic stack trace.
                var brokenPipe bool
                if ne, ok := err.(*net.OpError); ok {
                    if se, ok := ne.Err.(*os.SyscallError); ok {
                        if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
                            brokenPipe = true
                        }
                    }
                }

                httpRequest, _ := httputil.DumpRequest(c.Request, false)
                if brokenPipe {
                    zap.L().Error(c.Request.URL.Path,
                        zap.Any("error", err),
                        zap.String("request", string(httpRequest)),
                    )
                    // If the connection is dead, we can't write a status to it.
                    c.Error(err.(error)) // nolint: err check
                    c.Abort()
                    return
                }

                if stack {
                    zap.L().Error("[Recovery from panic]",
                        zap.Any("error", err),
                        zap.String("request", string(httpRequest)),
                        zap.String("stack", string(debug.Stack())),
                    )
                } else {
                    zap.L().Error("[Recovery from panic]",
                        zap.Any("error", err),
                        zap.String("request", string(httpRequest)),
                    )
                }
                c.AbortWithStatus(http.StatusInternalServerError)
            }
        }()
        c.Next()
    }
}
