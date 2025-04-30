package routes

import (
	"BlogService/db"
	"BlogService/modules/blog/controllers"
	repository "BlogService/modules/blog/repositories"
	"BlogService/modules/blog/services"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(router *gin.Engine) {

	userGroup := router.Group("/user")
	userControllers := controllers.NewUserController(services.NewUserService(repository.NewUserRepository(db.GetBlogDB())))
	{
		userGroup.GET("/", userControllers.GetUsers)
		userGroup.GET("/:id", userControllers.GetUserByID)
	}

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
