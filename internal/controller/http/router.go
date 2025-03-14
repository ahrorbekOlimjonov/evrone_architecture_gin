// Package v1 implements routing paths. Each services in own file.
package http

import (
	"net/http"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/gin-gonic/gin"
	// Swagger docs.
	_ "ai-seller/docs"
	"ai-seller/internal/controller/http/middleware"
	v1 "ai-seller/internal/controller/http/v1"
	"ai-seller/internal/usecase"
	"ai-seller/pkg/logger"
)

// NewRouter -.
// Swagger spec:
// @title       Go Clean Template API
// @description Using a translation service as an example
// @version     1.0
// @host        localhost:8080
// @BasePath    /v1
func NewRouter(app *gin.Engine, l logger.Interface, t usecase.UseCases) {
	// Options
	app.Use(middleware.Logger(l))
	app.Use(middleware.Recovery(l))

	// Swagger
	app.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// K8s probe
	app.GET("/healthz", func(ctx *gin.Context) { ctx.Status(http.StatusOK) })

	// Routers
	apiV1Group := app.Group("/v1")
	{
		// v1.NewAuthRoutes(apiV1Group, t, l)
		v1.NewProductRoutes(apiV1Group, t, l)
	}
}
