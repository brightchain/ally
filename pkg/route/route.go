package route

import (
	"ally/app/http/middleware"
	"ally/pkg/config"

	"github.com/gin-gonic/gin"
)

var route *gin.Engine

func SetRoute(r *gin.Engine) {
	route = r
}

func Run(r *gin.Engine) {
	r.Static("/public/storage", "./storage/app/public")
	r.Use(middleware.Session("SESSION_SECRET"))
	r.Run(config.Data.GetString("app.port"))
}
