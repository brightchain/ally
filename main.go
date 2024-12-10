package main

import (
	"h5/bootstrap"
	"h5/config"
	"h5/pkg/route"
	"os"

	"github.com/gin-gonic/gin"
)

func init() {
	config.Initialize()
}
func main() {
	bootstrap.SetupSlog()
	bootstrap.SetupModel()
	//bootstrap.SetupDatabase()
	gin.SetMode(os.Getenv("GIN_MODE"))
	r := bootstrap.SetupRoute()
	route.Run(r)
}
