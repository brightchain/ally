package main

import (
	"ally/config"
	"ally/model"
	"ally/routers"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	config.SetupSlog()
	config.InitSingleRedis()
	model.InitDb()
}

func main() {
	gin.SetMode(os.Getenv("GIN_MODE"))
	defer config.RedisClient.Close()
	routers.Run()
}
