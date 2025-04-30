package controllers

import (
	"BlogService/modules/blog/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userService services.UserService
}

func NewUserController(userService services.UserService) *UserController {
	return &UserController{userService: userService}
}

func (uc *UserController) GetUsers(ctx *gin.Context) {
	users, err := uc.userService.GetAllUsers()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}

func (uc *UserController) GetUserByID(ctx *gin.Context) {

	var id string

	id, has := ctx.Params.Get("id")
	if !has {
		id = ctx.Param("id")
		if len(id) == 0 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"err": "Param id cannt be empty!"})
			return
		}
	}

	user, err := uc.userService.GetUserByID(id)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, user)
}
