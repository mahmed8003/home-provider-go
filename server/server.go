package server

import (
	"home-provider/config"

	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/recover"
	"github.com/rs/zerolog"
)

type serverInfo struct {
	logger zerolog.Logger
	app    *iris.Application
}

var server *serverInfo

/*
Start :
*/
func Start(logger zerolog.Logger, config config.Server) {

	app := iris.New()

	// disbale default logger
	app.Logger().SetLevel("disable")

	// register recovery module
	app.Use(recover.New())

	// register request logger
	app.Use(NewLogger(logger, true, true, true, true, config.EnableLogs))

	// init routes
	initRoutes(app)

	// start http server
	app.Run(iris.Addr(config.Port))

	server = &serverInfo{
		logger: logger,
		app:    app,
	}
}

/*
GetServer :
*/
func GetServer() *iris.Application {
	return server.app
}
