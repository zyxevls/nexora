package routes

import (
	"nexora/internal/handler"
	"nexora/internal/middleware"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	api := r.Group("/api")

	api.POST("/register", handler.Register)
	api.POST("/login", handler.Login)

	api.GET("/products", handler.GetProducts)

	protected := api.Group("")
	protected.Use(middleware.AuthMiddleware())

	protected.GET("/cart", handler.GetCart)
	protected.POST("/cart/add", handler.AddToCart)
	protected.PUT("/cart/update", handler.UpdateCartItem)
	protected.DELETE("/cart/remove", handler.RemoveFromCart)
}
