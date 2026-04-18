package main

import (
	"log"

	"nexora/config"
	"nexora/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()
	config.InitDB()

	r := gin.Default()

	routes.SetupRoutes(r)

	log.Println("Nexora running on 8080")
	r.Run(":8080")
}
