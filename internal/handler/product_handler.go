package handler

import (
	"net/http"

	"nexora/internal/service"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	products := service.GetProducts()
	if products == nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"products": products})
}
