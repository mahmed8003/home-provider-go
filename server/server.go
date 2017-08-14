package server

import (
	"home-provider/app"

	"github.com/gin-gonic/gin"
)

/*
NewRouter :
*/
func NewRouter(ctx app.Context) *gin.Engine {

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

	return r
}

func errorHandler(c *gin.Context) {
	c.Next()

	// TODO: Handle it in a better way
	if len(c.Errors) > 0 {
		c.String(401, "I am here with error")
	}
}
