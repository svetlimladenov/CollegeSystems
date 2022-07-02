package main

import (
	"github.com/gin-gonic/gin"
	"github.com/svetlimladenov/collegesystems/pkg/college-server/controllers"
)

func main() {
	r := gin.New()

	// Middlewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Controllers
	userController := controllers.NewUserController(r)
	userController.ConfigureRoutes()

	r.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
