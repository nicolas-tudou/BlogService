package routes

import (
	"BlogService/controllers"
	"BlogService/db"
	repository "BlogService/repositories"
	"BlogService/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(router *gin.Engine) {

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "pong"})
	})

	userGroup := router.Group("/user")
	userControllers := controllers.NewUserController(services.NewUserService(repository.NewUserRepository(db.GetDB())))
	{
		userGroup.GET("/", userControllers.GetUsers)
		userGroup.GET("/:id", userControllers.GetUserByID)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
