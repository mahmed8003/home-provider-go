package main

import (
	"flag"
	"fmt"
	"home-provider/config"
	"home-provider/db"
	"home-provider/server"
	"home-provider/utils"
	"os"

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

	defer closeConnections()

	server.Start(logger, appConfig.Server)
}

func closeConnections() {
	db.DisconnectMongo()
	db.DisconnectRedis()
}
