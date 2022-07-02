package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	routerGroup gin.RouterGroup
}

func NewUserController(router *gin.Engine) UserController {
	controller := UserController{
		routerGroup: *router.Group("/users"),
	}

	return controller
}

func (controller *UserController) ConfigureRoutes() {
	controller.routerGroup.POST("/login", controller.Login)
	controller.routerGroup.POST("/register", controller.Register)
}

type LoginModel struct {
	Username string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (controller *UserController) Login(c *gin.Context) {
	var loginModel LoginModel

	if err := c.ShouldBindJSON(&loginModel); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "successful login")
}

func (controller *UserController) Register(c *gin.Context) {
	c.JSON(http.StatusOK, "Register Page")
}
