package main

import (
	"ally/routers"
	"ally/utils/logging"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	logging.SetupLog()
}

func main() {
	gin.SetMode(os.Getenv("GIN_MODE"))
	routers.Run()
}
