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
	model.InitDb()
}

func main() {
	gin.SetMode(os.Getenv("GIN_MODE"))
	routers.Run()
}
