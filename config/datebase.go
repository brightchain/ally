package config

import (
	"h5/pkg/config"
)

func init() {
	strMap := make(config.StrMap)
	for k := range config.GetStringMap("databases") {
		var value = map[string]interface{}{
			// 数据库连接信息
			"host":     config.Env("databases."+k+".host", "127.0.0.1"),
			"port":     config.Env("databases."+k+".port", "3306"),
			"database": config.Env("databases."+k+".database", "db"),
			"username": config.Env("databases."+k+".username", ""),
			"password": config.Env("databases."+k+".password", ""),
			"charset":  "utf8mb4",

			// 连接池配置
			"max_idle_connections": config.Env("databases."+k+".max_idle_connections", 25),
			"max_open_connections": config.Env("databases."+k+".max_open_connections", 100),
			"max_life_seconds":     config.Env("databases."+k+".max_life_seconds", 5*60),
		}
		strMap[k] = value
	}
	config.Add("databases", strMap)
}
