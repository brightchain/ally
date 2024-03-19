package config

import (
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"runtime"

	"gopkg.in/natefinch/lumberjack.v2"
)

var (
	DefaultPrefix      = ""
	DefaultCallerDepth = 0
	logPrefix          = ""
)

func SetupSlog() {
	var conf AppConfig
	err := GlobalConfig.Unmarshal(&conf)
	if err != nil {
		slog.Error("配置文件解析失败", err)
	}
	// 初始化日志
	logFile := &lumberjack.Logger{
		Filename:   conf.Logger.Filename,   // 日志文件的位置
		MaxSize:    conf.Logger.MaxSize,    // 文件最大尺寸（以MB为单位）
		MaxBackups: conf.Logger.MaxBackups, // 保留的最大旧文件数量
		MaxAge:     conf.Logger.MaxAge,     // 保留旧文件的最大天数
		Compress:   true,                   // 是否压缩/归档旧文件
		LocalTime:  true,                   // 使用本地时间创建时间戳
	}
	logOpts := slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}

	env := os.Getenv("GIN_MODE")
	if env == "release" {
		logOpts.Level = slog.LevelError
	}

	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(os.Stdout, logFile), &logOpts))

	slog.SetDefault(logger)
	logger.Info("初始化日志完成")

}

func setPrefix() string {
	_, file, line, ok := runtime.Caller(DefaultCallerDepth)
	if ok {
		logPrefix = fmt.Sprintf("[%s][%s:%d]", filepath.Base(file), line)
	} else {
		logPrefix = fmt.Sprintf("[%s]")
	}

	return logPrefix
}
