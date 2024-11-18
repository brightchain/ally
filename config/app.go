// Package config 应用的配置
package config

import "h5/pkg/config"

func init() {
	config.Add("app", config.StrMap{

		// 应用名称，暂时没有使用到
		"name": config.Env("app.name", "h5"),

		// 当前环境，用以区分多环境
		"env": config.Env("app.env", "production"),

		// 是否进入调试模式
		"debug": config.Env("app.debug", false),

		// 应用服务端口
		"port": config.Env("app.port", "3000"),

		// gorilla/sessions 在 Cookie 中加密数据时使用
		"key": config.Env("app.key", "33446a9dcf9ea060a0a6532b166da32f304af0de"),

		// age_128_key
		"aes_128_key": config.Env("crypto.aes_128.key", "1234567890123456"),
	})
}
