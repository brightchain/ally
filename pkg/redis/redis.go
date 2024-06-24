package redis

import (
	"ally/pkg/viperConf"
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/mitchellh/mapstructure"
	goRedis "github.com/redis/go-redis/v9"
)

var Client *goRedis.Client

func Initialize() {
	var redisConf viperConf.Redis
	rConf := viperConf.Data.Sub("redis")
	redisMap := rConf.AllSettings()
	mapstructure.Decode(redisMap, &redisConf)
	redisAddr := fmt.Sprintf("%s:%d", redisConf.Host, redisConf.Port)
	slog.Info("连接数据：%s", redisAddr)
	Client = goRedis.NewClient(&goRedis.Options{
		Addr:     redisAddr,
		Password: redisConf.Password,
		DB:       redisConf.Db,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	pong, err := Client.Ping(ctx).Result()
	if err != nil {
		slog.Info("redis connect failed, err:%v", err)
	} else {
		slog.Info("redis connect success, res:%v", pong)
	}

}
