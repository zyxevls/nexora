package routes

import (
	"nexora/internal/handler"
	"nexora/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")
	api.Use(middleware.AuthMiddleware())

	api.POST("/register", handler.Register)
	api.POST("/login", handler.Login)

	api.GET("/products", handler.GetProducts)
}
