package controllers

import "github.com/gin-gonic/gin"

type BaseController struct {
}

func (c *BaseController) Success(ctx *gin.Context) {
	ctx.String(200, "成功")
}

func (c *BaseController) Fail(ctx *gin.Context) {
	ctx.String(200, "失败")
}
