package api

import "github.com/gin-gonic/gin"

func Index(c *gin.Context) {
	str := c.Query("decrypt")
	c.JSON(200, gin.H{
		"status": "success",
		"data":   str,
	})
}
