package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/svetlimladenov/collegesystems/pkg/models"
	"github.com/svetlimladenov/collegesystems/pkg/services"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController() UserController {
	return UserController{}
}

func (controller *UserController) Login(c *gin.Context) {
	var loginModel models.LoginInputModel

	if err := c.ShouldBindJSON(&loginModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	controller.UserService.Login(loginModel.Username, loginModel.Password)

	c.JSON(http.StatusOK, "successful login")
}

func (controller *UserController) RegisterAction(c *gin.Context) {
	var model models.RegisterInputModel

	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := controller.UserService.Register(model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, "User Registered")
}
