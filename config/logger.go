package config

import (
	"io"
	"log/slog"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

func SetupSlog() {
	logConf := GlobalConfig.Sub("logger")
	filename := logConf.GetString("filename")
	maxSize := logConf.GetInt("maxSize")
	maxBackups := logConf.GetInt("maxBackups")
	maxAge := logConf.GetInt("maxAge")
	level := logConf.GetString("level")
	// 初始化日志
	logFile := &lumberjack.Logger{
		Filename:   filename,   // 日志文件的位置
		MaxSize:    maxSize,    // 文件最大尺寸（以MB为单位）
		MaxBackups: maxBackups, // 保留的最大旧文件数量
		MaxAge:     maxAge,     // 保留旧文件的最大天数
		Compress:   true,       // 是否压缩/归档旧文件
		LocalTime:  true,       // 使用本地时间创建时间戳
	}
	logOpts := slog.HandlerOptions{
		AddSource: true,
		Level:     slog.LevelDebug,
	}

	switch level {
	case "debug":
		logOpts.Level = slog.LevelDebug
	case "info":
		logOpts.Level = slog.LevelInfo
	case "error":
		logOpts.Level = slog.LevelError
	case "warn":
		logOpts.Level = slog.LevelWarn
	default:
		logOpts.Level = slog.LevelDebug
	}

	logger := slog.New(slog.NewJSONHandler(io.MultiWriter(os.Stdout, logFile), &logOpts))

	slog.SetDefault(logger)
	logger.Info("初始化日志完成")

}
