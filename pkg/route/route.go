package route

import (
	"ally/middleware"
	"ally/pkg/viperConf"

	"github.com/gin-gonic/gin"
)

var route *gin.Engine

func SetRoute(r *gin.Engine) {
	route = r
}

func Run(r *gin.Engine) {
	r.Static("/public/storage", "./storage/app/public")
	r.Use(middleware.Session("SESSION_SECRET"))
	r.Run(viperConf.Data.GetString("app.port"))
}
