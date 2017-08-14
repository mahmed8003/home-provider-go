package server

import (
	"home-provider/app"

	routing "github.com/go-ozzo/ozzo-routing"
)

/*
NewRouter :
*/
func NewRouter(ctx app.Context) *routing.Router {

	config := ctx.Config().Server
	router := routing.New()
	router.IgnoreTrailingSlash = true

	router.Use(
		NewRequestLogger(ctx.Logger(), config.EnableLogs),
	)

	addRoutes(ctx, router)

	return router
}
