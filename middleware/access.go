package middleware

import "github.com/gin-gonic/gin"

func ExportExport() gin.HandlerFunc {
	return func(c *gin.Context) {
		at := c.Query("at")
		if at != "sfdjwie2ji239324" {
			c.String(200, "非法访问！")
			c.Abort()
			return
		}
		c.Next()

	}
}
