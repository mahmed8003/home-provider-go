package server

import (
	"home-provider/controllers"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

func initRoutes(app *iris.Application) {
	app.Get("/ping", func(ctx context.Context) {
		ctx.WriteString("pong")
	})

	v1 := app.Party("/v1")
	{
		// init user controller
		user := new(controllers.UserController)

		// create user routes
		usersRoutes := v1.Party("/users")
		usersRoutes.Get("/", user.CreateUser)
	}
}
