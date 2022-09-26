package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/svetlimladenov/collegesystems/models"
	"github.com/svetlimladenov/collegesystems/services"
)

type LoginInputModel struct {
	Username string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func Login(c *gin.Context) {
	var loginModel LoginInputModel

	if err := c.ShouldBindJSON(&loginModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	us := services.UserService{}
	if err := us.Login(loginModel.Username, loginModel.Password); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "successful login")
}

type RegisterInputModel struct {
	Username  string `form:"username" json:"username" binding:"required"`
	Password  string `form:"password" json:"password" binding:"required"`
	FirstName string `form:"firstName" json:"firstName" binding:"required"`
	LastName  string `form:"lastName" json:"lastName" binding:"required"`
	Email     string `form:"email" json:"email" binding:"required"`
}

func Register(c *gin.Context) {
	var model RegisterInputModel

	if err := c.ShouldBindJSON(&model); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Username:  model.Username,
		Password:  model.Password,
		FirstName: model.FirstName,
		LastName:  model.LastName,
		Email:     model.Email,
	}

	us := services.UserService{}

	if err := us.Register(user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusAccepted, "User Registered")
}
