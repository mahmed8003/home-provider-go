package server

import (
	"home-provider/config"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type serverInfo struct {
	logger zerolog.Logger
	engine *gin.Engine
}

var server *serverInfo

/*
Start :
*/
func Start(logger zerolog.Logger, config config.Server) {

	// Creates a router without any middleware by default
	r := gin.New()

	// Global middleware
	//r.Use(gin.Logger())
	r.Use(Logger(logger, config.EnableLogs))
	r.Use(gin.Recovery())

	// init routes
	initRoutes(r)

	// Listen and serve
	r.Run(config.Port)

	server = &serverInfo{
		logger: logger,
		engine: r,
	}
}

/*
GetServer :
*/
func GetServer() *gin.Engine {
	return server.engine
}
