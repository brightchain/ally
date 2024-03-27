package routers

import (
	"ally/config"
	"ally/middleware"

	"github.com/gin-gonic/gin"
)

var r = gin.Default()

// Run will start the server
func Run() {
	r.Static("/public/storage", "./storage/app/public")
	r.Use(middleware.Session("SESSION_SECRET"))
	getRoutes()
	r.Run(config.GlobalConfig.GetString("app.port"))
}

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func getRoutes() {
	addWebRoute(r)
	addApiRoute(r)
}
