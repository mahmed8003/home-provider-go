package server

import (
	"home-provider/app"
	"home-provider/controllers"

	routing "github.com/go-ozzo/ozzo-routing"
	"github.com/go-ozzo/ozzo-routing/content"
)

func addRoutes(ctx app.Context, router *routing.Router) {

	v1 := router.Group("/v1")
	v1.Use(content.TypeNegotiator(content.JSON))
	{
		userRouter := v1.Group("/users")
		{
			user := controllers.NewUserController(ctx)
			userRouter.Get("", user.CreateUser)
		}
	}
}
