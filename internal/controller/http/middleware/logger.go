package middleware

import (
	"strings"

	"ai-seller/pkg/logger"

	"github.com/gin-gonic/gin"
)

func buildRequestMessage(ctx *gin.Context) string {
	var result strings.Builder

	result.WriteString(ctx.ClientIP())
	result.WriteString(" - ")
	result.WriteString(ctx.Request.Method)
	result.WriteString(" ")
	result.WriteString(ctx.Request.RequestURI)
	result.WriteString(" ")

	return result.String()
}

func Logger(l logger.Interface) func(c *gin.Context) {
	return func(ctx *gin.Context) {
		ctx.Next()
		l.Info(buildRequestMessage(ctx))
	}
}
