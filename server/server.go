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

	// add routes
	addRoutes(ctx, r)

	return r
}
