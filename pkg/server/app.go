package server

import (
	"github.com/gin-gonic/gin"
	"github.com/svetlimladenov/collegesystems/pkg/server/controllers"
)

type App struct {
	Router *gin.Engine
}

func NewApp() App {
	app := App{
		Router: gin.New(),
	}

	user := app.Router.Group("/user")

	user.Handle("POST", "/login", InjectController)
	user.Handle("POST", "/register", Register)

	return app
}

func InjectController(ctx *gin.Context) {
	controller := controllers.NewUserController()
	controller.Login(ctx)
}

func Register(ctx *gin.Context) {
	controller := controllers.NewUserController()
	controller.RegisterAction(ctx)
}

func (app *App) Run() {
	app.Router.Run()
}
