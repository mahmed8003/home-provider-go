package db

import (
	"home-provider/config"

	"github.com/go-redis/redis"
	"github.com/rs/zerolog"
)

type redisConnecion struct {
	logger zerolog.Logger
	client *redis.Client
}

var redisCon *redisConnecion

/*
ConnectRedis :
*/
func ConnectRedis(logger zerolog.Logger, config config.Redis) error {
	logger.Info().Msg("Connecting to redis")
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})

	redisCon = &redisConnecion{
		logger: logger,
		client: client,
	}
	_, err := client.Ping().Result()
	return err
}

/*
GetRedis :
*/
func GetRedis() *redis.Client {
	return redisCon.client
}

/*
DisconnectRedis :
*/
func DisconnectRedis() {
	if redisCon != nil {
		redisCon.logger.Info().Msg("Disconnecting from redis")
		redisCon.client.Close()
	}
}
