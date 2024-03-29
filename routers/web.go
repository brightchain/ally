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
		export.GET("/ydln", controllers.Ydln)
		export.GET("/excel-down", controllers.ExcelDown)
	}

	aes := controllers.AesEcb{}
	r.POST("/aes/aes", aes.Aes)
	r.POST("/aes/encrypt", aes.Encrypt)
	r.POST("/aes/dow", aes.Down)

}
