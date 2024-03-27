package routers

import (
	"ally/api"
	"ally/middleware"

	"github.com/gin-gonic/gin"
)

func addApiRoute(r *gin.Engine) {
	r.POST("/zip", api.Zip)
	apiGroup := r.Group("/api")
	apiGroup.Use(middleware.AesDecrypt())
	{
		apiGroup.POST("/downzip", api.PhotoOrder)

	}
}
