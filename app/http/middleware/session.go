package middleware

import (
	"h5/pkg/session"

	"github.com/gin-gonic/gin"
)

func StartSession() gin.HandlerFunc {
	return func(c *gin.Context) {
		session.StartSession(c.Writer, c.Request)
		c.Next()
	}
}
