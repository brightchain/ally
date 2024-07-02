package config

import "ally/pkg/config"

func init() {
	config.Add("logger", config.StrMap{

		//日志文件路径
		"name": config.Env("LOGGER_NAME", "ally"),

		// 文件最大尺寸（以MB为单位）
		"size": config.Env("LOGGER_SIZE", 4),

		// 保留的最大旧文件数量
		"maxBackups": config.Env("LOGGER_MAX_BACKUPS", 10),

		// 保留旧文件的最大天数
		"maxAge": config.Env("LOGGER_MAX_AGE", 30),

		// 日志模式
		"level": config.Env("LOGGER_LEVEL", "debug"),
	})
}
