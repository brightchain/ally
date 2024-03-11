package main

import (
	"h5/routers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	routers.WebRoute(r)
	// Listen and Server in 0.0.0.0:8080
	r.Run("localhost:8787")
}
