package bootstrap

import "h5/pkg/goredis"

func SetupRedis() {
	goredis.Initialize()
}
