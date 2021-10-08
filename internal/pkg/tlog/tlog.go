package tlog

import (
	"fmt"
	"io"
	"os"
	"path"
	"time"

	"htblog/internal/pkg/utils"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	*zap.SugaredLogger
}

var (
	_defaultLogger *Logger
	_logPath       string
	_logName       string
)

func Init(outPath []string) {
	logPath := os.Getenv("LOG_PATH")
	if logPath == "" {
		logPath = "var/log"
	}
	if ok, _ := utils.PathExists(logPath); !ok { // 判断是否有Director文件夹
		fmt.Printf("create %v directory\n", logPath)
		_ = os.Mkdir(logPath, os.ModePerm)
	}
	_logPath = logPath
	_logName = "htblog.log"

	pe := zap.NewDevelopmentEncoderConfig()
	pe.EncodeTime = func(t time.Time, encoder zapcore.PrimitiveArrayEncoder) {
		encoder.AppendString(t.Format("2006-01-02 15:04:05"))
	}
	pe.ConsoleSeparator = " | "

	consoleEncoder := zapcore.NewConsoleEncoder(pe)
	zap.NewDevelopment()

	level := zap.DebugLevel

	cores := make([]zapcore.Core, 0, len(outPath))
	for _, path := range outPath {
		switch path {
		case "console":
			cores = append(cores, zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), level))
		case "file":
			cores = append(cores, zapcore.NewCore(consoleEncoder, zapcore.AddSync(getFileWriter()), level))
		}
	}

	core := zapcore.NewTee(cores...)

	l := zap.New(core, zap.AddCaller())
	_defaultLogger = &Logger{
		SugaredLogger: l.Sugar(),
	}

	fmt.Println("init log completed.")
}

func Tag(tag string) *Logger {
	return &Logger{
		SugaredLogger: _defaultLogger.Named(tag),
	}
}

func getFileWriter() io.Writer {
	logPath := path.Join(_logPath, _logName)
	fmt.Println("init log file:", logPath)

	flog := &lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    500, // megabytes
		MaxBackups: 3,
		MaxAge:     15,    //days
		Compress:   false, // disabled by default
	}
	return flog
}
