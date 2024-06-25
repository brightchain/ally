package routers

import (
	"ally/api"
	"ally/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterApiRouters(r *gin.Engine) {
	r.POST("/zip", api.Zip)
	r.POST("/redis", api.Redis)
	apiGroup := r.Group("/api")
	apiGroup.Use(middleware.AesDecrypt())
	{
		apiGroup.POST("/downzip", api.PhotoOrderCy)
	}
}
