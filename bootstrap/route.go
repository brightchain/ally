package bootstrap

import (
	"h5/pkg/route"
	"h5/routers"

	"github.com/gin-gonic/gin"
)

func SetupRoute() *gin.Engine {
	r := gin.Default()
	routers.RegisterWebRouters(r)
	routers.RegisterApiRouters(r)
	route.SetRoute(r)
	return r
}
