package config

import "h5/pkg/config"

func init() {
	config.Add("redis", config.StrMap{

		// 地址
		"host": config.Env("redis.host", "127.0.0.1"),

		// 端口
		"port": config.Env("redis.prot", 6379),

		// redis密码
		"password": config.Env("redis.password", ""),

		// 数据库
		"db": config.Env("redis.db", 1),
	})
}
