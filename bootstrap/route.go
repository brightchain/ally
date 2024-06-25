package bootstrap

import (
	"ally/pkg/route"
	"ally/routers"

	"github.com/gin-gonic/gin"
)

func SetupRoute() *gin.Engine {
	r := gin.Default()
	routers.RegisterWebRouters(r)
	routers.RegisterApiRouters(r)
	route.SetRoute(r)
	return r
}
