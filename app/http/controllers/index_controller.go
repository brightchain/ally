package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Index struct{}

func (*Index) Index(c *gin.Context) {
	c.String(http.StatusOK, "测试页面")
}
