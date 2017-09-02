package main

import (
	"context"
	"flag"
	"fmt"
	"home-provider/app"
	"home-provider/config"
	"home-provider/db"
	"home-provider/server"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"go.uber.org/zap"
)

func main() {
	enviroment := flag.String("e", "development", "")
	flag.Usage = func() {
		fmt.Println("Usage: main -e {mode}")
		os.Exit(1)
	}
	flag.Parse()

	appConfig, err := config.LoadConfig(*enviroment)
	if err != nil {
		panic(fmt.Errorf("Invalid application configuration: %s", err))
	}

	logger, _ := zap.NewDevelopment()
	defer logger.Sync()

	database, err := db.ConnectMongo(logger, appConfig.Database)
	if err != nil {
		logger.Fatal("Failed to connec to database", zap.Error(err))
	}

	redis, err := db.ConnectRedis(logger, appConfig.Redis)
	if err != nil {
		logger.Fatal("Failed to connec to redis", zap.Error(err))
	}

	addr := ":" + os.Getenv("PORT")
	if addr == ":" {
		addr = appConfig.Server.Port
	}
	appConfig.Server.Port = addr
	appCtx := app.NewContext(*enviroment, appConfig, logger, redis, database)

	// create router
	router := server.NewRouter(appCtx)

	// create http server
	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}
	go func() {
		logger.Info("Server listening at " + addr)
		if err := server.ListenAndServe(); err != nil {
			logger.Error("Server error", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM)
	signal.Notify(quit, syscall.SIGINT)
	<-quit

	logger.Info("Shutting down server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		logger.Fatal("Server shutdown", zap.Error(err))
	}
	database.Close()
	redis.Close()
	logger.Info("Exiting ...")
}
