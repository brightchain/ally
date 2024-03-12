package routers

import "github.com/gin-gonic/gin"

var router = gin.Default()

// Run will start the server
func Run() {
	getRoutes()
	router.Run(":8787")
}

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func getRoutes() {
	addWebRoute(router)
}
