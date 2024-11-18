package config

import "h5/pkg/config"

func init() {
	config.Add("logger", config.StrMap{

		//日志文件路径
		"name": config.Env("logger.filename", "./app.log"),

		// 文件最大尺寸（以MB为单位）
		"size": config.Env("logger.maxSize", 4),

		// 保留的最大旧文件数量
		"maxBackups": config.Env("logger.maxBackups", 10),

		// 保留旧文件的最大天数
		"maxAge": config.Env("logger.maxAge", 30),

		// 日志模式
		"level": config.Env("logger.level", "debug"),

		//gorm日志文件路径
		"gormName": config.Env("logger.gormFile", "./gorm.log"),
	})
}
