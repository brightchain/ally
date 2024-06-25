package main

import (
	"ally/bootstrap"
	"ally/pkg/route"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	bootstrap.SetupViper()
	bootstrap.SetupSlog()
	bootstrap.SetupModel()
	//bootstrap.SetupDatabase()
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := bootstrap.SetupRoute()
	route.Run(r)
}
