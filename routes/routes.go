package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/getusers", func(context *gin.Context) {
		context.JSON(http.StatusOK, "success")
	})
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
