package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.GET("/getusers", func(context *gin.Context) {
		context.JSON(http.StatusOK, "success")
	})
}
