package server

import (
	"home-provider/config"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type serverInfo struct {
	logger zerolog.Logger
	router *gin.Engine
}

var server *serverInfo

/*
Start :
*/
func NewRouter(logger zerolog.Logger, config config.Server, env string) *gin.Engine {

	// Creates a router
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	// Global middleware
	r.Use(Logger(logger, config.EnableLogs))
	r.Use(gin.Recovery())

	// init routes
	initRoutes(r)

	// Listen and serve
	//r.Run(config.Port)

	server = &serverInfo{
		logger: logger,
		router: r,
	}

	return r
}

/*
GetRouter :
*/
func GetRouter() *gin.Engine {
	return server.router
}
