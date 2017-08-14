package server

import (
	"home-provider/app"

	"github.com/gin-gonic/gin"
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
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	addRoutes(ctx, r)

	/*

		// Creates a router
		if ctx.Env() == "production" {
			gin.SetMode(gin.ReleaseMode)
		}

		r := gin.New()

		config := ctx.Config().Server
		// Global middleware
		r.Use(Logger(ctx.Logger(), config.EnableLogs))
		r.Use(gin.Recovery())
		r.Use(errorHandler)

		// add routes
		addRoutes(ctx, r)
	*/

	return r
}

func errorHandler(c *gin.Context) {
	c.Next()

	// TODO: Handle it in a better way
	if len(c.Errors) > 0 {
		c.String(401, "I am here with error")
	}
}
