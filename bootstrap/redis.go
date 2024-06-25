package bootstrap

import "ally/pkg/goredis"

func SetupRedis() {
	goredis.Initialize()
}
