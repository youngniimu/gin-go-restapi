package controllers

import (
	"go-rest-api/models"
	"go-rest-api/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func New(userservice services.UserService) UserController {
	return UserController{
		UserService: userservice,
	}
}

func (uc *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := uc.UserService.CreateUser(&user); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusCreated, &user)
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	user, err := uc.UserService.GetUser(ctx.Params.ByName("personalCode"))
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &user)
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	allUsers, err := uc.UserService.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &allUsers)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := uc.UserService.UpdateUser(&user, ctx.Params.ByName("personalCode")); err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, &user)
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	if err := uc.UserService.DeleteUser(ctx.Params.ByName("personalCode")); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, nil)

}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user")
	userRoute.POST("", uc.CreateUser)
	userRoute.GET("/:personalCode", uc.GetUser)
	userRoute.GET("", uc.GetAll)
	userRoute.PATCH("/:personalCode", uc.UpdateUser)
	userRoute.DELETE("/:personalCode", uc.DeleteUser)
}
