package main

import (
	"ally/bootstrap"
	"ally/routers"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	bootstrap.SetupViper()
	bootstrap.SetupSlog()
	bootstrap.SetupDatabase()
	gin.SetMode(os.Getenv("GIN_MODE"))
	routers.Run()
}
