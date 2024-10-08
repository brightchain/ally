package routers

import (
	"ally/app/http/controllers"
	"ally/app/http/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterWebRouters(r *gin.Engine) {
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	i := new(controllers.Index)
	r.GET("/index", i.Index)
	dir := new(controllers.DirectoryClear)
	r.GET("/photo-clear", dir.PhotoDirClear)
	r.GET("/album-clear", dir.AlbumDirClear)
	r.GET("/calendar-clear", dir.CalendarDirClear)
	r.GET("/tshirt-clear", dir.TshirtDirClear)

	export := new(controllers.ExportExcel)
	r.GET("/hngx", export.Hngx)
	r.GET("/hnkj", export.Hnkj)
	r.GET("/wj", export.Smwj)
	exGroup := r.Group("/export")
	exGroup.Use(middleware.ExportExport())
	{
		exGroup.GET("/fjpa", export.Fjpa)
		exGroup.GET("/xinhua", export.Xinhua)
		exGroup.GET("/ydln", export.Ydln)
		exGroup.GET("/shtp", export.ShTp)
		exGroup.GET("/fjtp", export.FjTp)

	}

	aes := controllers.AesEcb{}
	r.POST("/aes/aes", aes.Aes)
	r.POST("/aes/encrypt", aes.Encrypt)
	r.POST("/aes/dow", aes.Down)

	order := controllers.PayOrder{}
	orderGroup := r.Group("/order")
	orderGroup.Use(middleware.ExportExport())
	{
		orderGroup.GET("/product", order.GetOrderProduct)

	}

}
