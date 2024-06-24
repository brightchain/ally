package bootstrap

import "ally/pkg/redis"

func SetupRedis() {
	redis.Initialize()
}
