package tlog

import (
	"fmt"
	"htblog/internal/pkg/tconfig"
	"os"
	"time"

	"htblog/internal/pkg/utils"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var level zapcore.Level

func init() {
	logPath := os.Getenv("LOG_PATH")
	if logPath == "" {
		logPath = "var/log"
	}
	if ok, _ := utils.PathExists(logPath); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", logPath)
		_ = os.Mkdir(logPath, os.ModePerm)
	}
}

func Tag(tag string) *zap.Logger {
	logPath := tconfig.GetString("log.path")
	if logPath == "" {
		logPath = "var/log"
	}
	if ok, _ := utils.PathExists(logPath); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", logPath)
		_ = os.Mkdir(logPath, os.ModePerm)
	}

	var level zapcore.Level
	switch tconfig.Get("log.level") { // 初始化配置文件的Level
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	case "dpanic":
		level = zap.DPanicLevel
	case "panic":
		level = zap.PanicLevel
	case "fatal":
		level = zap.FatalLevel
	default:
		level = zap.InfoLevel
	}

	var logger *zap.Logger
	if level == zap.DebugLevel || level == zap.ErrorLevel {
		logger = zap.New(getEncoderCore(), zap.AddStacktrace(level))
	} else {
		logger = zap.New(getEncoderCore())
	}
	logger = logger.WithOptions(zap.AddCaller())
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
		StacktraceKey:  "string",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseColorLevelEncoder,
		EncodeTime:     CustomTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.FullCallerEncoder,
	}
	return config
}

// getEncoder 获取zapcore.Encoder
func getEncoder() zapcore.Encoder {
	if global.GVA_CONFIG.Zap.Format == "json" {
		return zapcore.NewJSONEncoder(getEncoderConfig())
	}
	return zapcore.NewConsoleEncoder(getEncoderConfig())
}

// getEncoderCore 获取Encoder的zapcore.Core
func getEncoderCore() (core zapcore.Core) {
	writer, err := utils.GetWriteSyncer() // 使用file-rotatelogs进行日志分割
	if err != nil {
		fmt.Printf("Get Write Syncer Failed err:%v", err.Error())
		return
	}
	return zapcore.NewCore(getEncoder(), writer, level)
}

// 自定义日志输出时间格式
func CustomTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format(global.GVA_CONFIG.Zap.Prefix + "2006/01/02 - 15:04:05.000"))
}
