package routers

import (
	"h5/controllers"

	"github.com/gin-gonic/gin"
)

func addWebRoute(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/index", controllers.Index)
	r.GET("/xinhua", controllers.Xinhua)
	r.GET("/hngx", controllers.Hngx)
	r.GET("/hnkj", controllers.Hnkj)
	r.GET("/wj", controllers.Smwj)

}
