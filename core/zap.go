/**
    @Author:     ZonzeeLi
    @Project:    spikesystem
    @CreateDate: 12/21/2021
    @UpdateDate: 12/22/2021
    @Note:       zap日志自定义配置
**/

package core

import (
	"fmt"
	"os"
	"spikesystem/global"
	"spikesystem/utils"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Zap() (logger *zap.Logger) {
	if ok, _ := utils.PathExists(global.Spike_CONFIG.Zap.Director); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", global.Spike_CONFIG.Zap.Director)
		_ = os.Mkdir(global.Spike_CONFIG.Zap.Director, os.ModePerm)
	}
	// 调试级别
	debugPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.DebugLevel
	})
	// 日志级别
	infoPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.InfoLevel
	})
	// 警告级别
	warnPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev == zap.WarnLevel
	})
	// 错误级别
	errorPriority := zap.LevelEnablerFunc(func(lev zapcore.Level) bool {
		return lev >= zap.ErrorLevel
	})

	cores := [...]zapcore.Core{
		getEncoderCore(fmt.Sprintf("./%s/server_debug.log", global.Spike_CONFIG.Zap.Director), debugPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_info.log", global.Spike_CONFIG.Zap.Director), infoPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_warn.log", global.Spike_CONFIG.Zap.Director), warnPriority),
		getEncoderCore(fmt.Sprintf("./%s/server_error.log", global.Spike_CONFIG.Zap.Director), errorPriority),
	}
	logger = zap.New(zapcore.NewTee(cores[:]...), zap.AddCaller())

	if global.Spike_CONFIG.Zap.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

// getEncoderConfig 获取zapcore.EncoderConfig
func getEncoderConfig() (config zapcore.EncoderConfig) {
	config = zapcore.EncoderConfig{
		MessageKey:     "message",
		LevelKey:       "level",
		TimeKey:        "time",
		NameKey:        "logger",
		CallerKey:      "caller",
		StacktraceKey:  global.Spike_CONFIG.Zap.StacktraceKey,
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	switch {
	case global.Spike_CONFIG.Zap.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	case global.Spike_CONFIG.Zap.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
	case global.Spike_CONFIG.Zap.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		config.EncodeLevel = zapcore.CapitalLevelEncoder
	case global.Spike_CONFIG.Zap.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
	default:
		config.EncodeLevel = zapcore.LowercaseLevelEncoder
	}
	return config
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	if global.Spike_CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore(fileName string, level zapcore.LevelEnabler) (core zapcore.Core) {
	writer := GetWriteSyncer(fileName) // 使用file-rotatelogs进行日志分割
	return zapcore.NewCore(getEncoder(), writer, level)
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.Spike_CONFIG.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
}

func GetWriteSyncer(file string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   file, // 日志文件的位置
		MaxSize:    10,   // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 200,  // 保留旧文件的最大个数
		MaxAge:     30,   // 保留旧文件的最大天数
		Compress:   true, // 是否压缩/归档旧文件
	}

	if global.Spike_CONFIG.Zap.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	}
	return zapcore.AddSync(lumberJackLogger)
}

//// GinLogger 接收gin框架默认的日志
//func GinLogger() gin.HandlerFunc {
//	return func(c *gin.Context) {
//		start := time.Now()
//		path := c.Request.URL.Path
//		query := c.Request.URL.RawQuery
//		c.Next()
//
//		cost := time.Since(start)
//		zap.L().Info(path,
//			zap.Int("status", c.Writer.Status()),
//			zap.String("method", c.Request.Method),
//			zap.String("path", path),
//			zap.String("query", query),
//			zap.String("ip", c.ClientIP()),
//			zap.String("user-agent", c.Request.UserAgent()),
//			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
//			zap.Duration("cost", cost),
//		)
//	}
//}
//
//// GinRecovery recover掉项目可能出现的panic，并使用zap记录相关日志
//func GinRecovery(stack bool) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		defer func() {
//			if err := recover(); err != nil {
//				// Check for a broken connection, as it is not really a
//				// condition that warrants a panic stack trace.
//				var brokenPipe bool
//				if ne, ok := err.(*net.OpError); ok {
//					if se, ok := ne.Err.(*os.SyscallError); ok {
//						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
//							brokenPipe = true
//						}
//					}
//				}
//
//				httpRequest, _ := httputil.DumpRequest(c.Request, false)
//				if brokenPipe {
//					zap.L().Error(c.Request.URL.Path,
//						zap.Any("error", err),
//						zap.String("request", string(httpRequest)),
//					)
//					// If the connection is dead, we can't write a status to it.
//					c.Error(err.(error)) // nolint: errcheck
//					c.Abort()
//					return
//				}
//
//				if stack {
//					zap.L().Error("[Recovery from panic]",
//						zap.Any("error", err),
//						zap.String("request", string(httpRequest)),
//						zap.String("stack", string(debug.Stack())),
//					)
//				} else {
//					zap.L().Error("[Recovery from panic]",
//						zap.Any("error", err),
//						zap.String("request", string(httpRequest)),
//					)
//				}
//				c.AbortWithStatus(http.StatusInternalServerError)
//			}
//		}()
//		c.Next()
//	}
//}
