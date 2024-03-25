package routers

import (
	"ally/api"
	"ally/middleware"

	"github.com/gin-gonic/gin"
)

func addApiRoute(r *gin.Engine) {
	apiGroup := r.Group("/api")
	apiGroup.Use(middleware.AesDecrypt())
	{
		apiGroup.POST("/downZip", api.Index)
	}
}
