package main

import (
	"ally/config"
	"ally/routers"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	config.SetupSlog()
}

func main() {
	gin.SetMode(os.Getenv("GIN_MODE"))
	routers.Run()
}
