package server

import (
	"home-provider/app"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

/*
NewRouter :
*/
func NewRouter(ctx app.Context) chi.Router {

	r := chi.NewRouter()

	// A good base middleware stack
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	//r.Use(middleware.Logger)
	r.Use(NewRequestLogger(ctx.Logger(), true))
	r.Use(middleware.Recoverer)

	addRoutes(ctx, r)

	return r
}
