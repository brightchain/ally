package goredis

import (
	"context"
	"fmt"
	"h5/pkg/config"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

var Client *redis.Client

func Initialize() {

	Client = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%v:%v", config.GetString("redis.host"), config.GetString("redis.port")),
		Password: config.GetString("redis.password"),
		DB:       config.GetInt("redis.db"),
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
