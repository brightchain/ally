package config

import "ally/pkg/config"

func init() {
	config.Add("redis", config.StrMap{

		// 地址
		"host": config.Env("REDIS_HOST", "127.0.0.1"),

		// 端口
		"port": config.Env("REDIS_PORT", 6379),

		// redis密码
		"password": config.Env("REDIS_PASSWORD", ""),
		
		// 数据库
		"db": config.Env("REDIS_DB", 1),


	})
}
