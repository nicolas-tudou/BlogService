package main

import (
	"BlogService/config"
	"BlogService/middlewares"
	"BlogService/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {

	config.LoadConfig()

	if config.AppConfig.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.Default()

	router.Use(
		gin.Logger(),
		gin.Recovery(),
		middlewares.JWTAuth(),
	)

	routes.RegisterRoutes(router)

	log.Printf("Service is running in port: %s\n", config.AppConfig.Port)

	if err := router.Run(":" + config.AppConfig.Port); err != nil {
		log.Fatal("Failed to start service: %v\n", err)
	}
}
