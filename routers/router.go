package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/svetlimladenov/collegesystems/routers/api"
)

func InitRouter() http.Handler {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	user := r.Group("/user")

	user.POST("/login", api.Login)
	user.POST("/register", api.Register)

	return r
}
