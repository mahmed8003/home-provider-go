package server

import (
	"home-provider/config"

	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
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
	app.Logger().SetLevel("disable")

	app.Use(recover.New())
	app.Use(NewLogger(logger, true, true, true, true, config.EnableLogs))

	app.Get("/ping", func(ctx context.Context) {
		ctx.WriteString("pong")
	})

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
