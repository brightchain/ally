package routers

import (
	"ally/controllers"
	"ally/middleware"

	"github.com/gin-gonic/gin"
)

func addWebRoute(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	r.GET("/index", controllers.Index)
	r.GET("/hngx", controllers.Hngx)
	r.GET("/hnkj", controllers.Hnkj)
	r.GET("/wj", controllers.Smwj)

	export := r.Group("/export")
	export.Use(middleware.ExportExport())
	{
		export.GET("/fjpa", controllers.Fjpa)
		export.GET("/xinhua", controllers.Xinhua)
	}

}
