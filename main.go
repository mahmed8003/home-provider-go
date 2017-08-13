package main

import (
	"context"
	"flag"
	"fmt"
	"home-provider/config"
	"home-provider/db"
	"home-provider/server"
	"home-provider/utils"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
)

func main() {
	enviroment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: main -e {mode}")
		os.Exit(1)
	}
	flag.Parse()

	if err := config.LoadConfig(*enviroment); err != nil {
		panic(fmt.Errorf("Invalid application configuration: %s", err))
	}

	appConfig := config.GetConfig()

	zerolog.SetGlobalLevel(utils.GetLogLevelByString(appConfig.LogLevel))
	logger := zerolog.New(os.Stdout).With().Timestamp().Logger()

	if err := db.ConnectMongo(logger, appConfig.Database); err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to database")
	} else {
		logger.Info().Msg("Database connection successfull")
	}

	if err := db.ConnectRedis(logger, appConfig.Redis); err != nil {
		logger.Fatal().Err(err).Msg("Failed to connect to redis")
	} else {
		logger.Info().Msg("Redis connection successfull")
	}

	addr := ":" + os.Getenv("PORT")
	if addr == ":" {
		addr = appConfig.Server.Port
	}

	// create router
	router := server.NewRouter(logger, appConfig.Server, *enviroment)

	// create http server
	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	go func() {
		logger.Info().Msg("Server listening at " + addr)
		if err := server.ListenAndServe(); err != nil {
			logger.Error().Err(err).Msg("Server error")
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM)
	signal.Notify(quit, syscall.SIGINT)
	<-quit

	logger.Info().Msg("Shutting down server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal().Err(err).Msg("Server shutdown")
	}
	db.DisconnectMongo()
	db.DisconnectRedis()
	logger.Info().Msg("Exiting ...")
}
