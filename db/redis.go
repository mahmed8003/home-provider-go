package db

import (
	"home-provider/config"

	"go.uber.org/zap"

	"github.com/go-redis/redis"
)

/*
Redis :
*/
type Redis interface {

	// GetClient return redis client
	GetClient() *redis.Client

	// Close the redis, freeing up any available resources.
	Close()
}

type redisConnecion struct {
	logger *zap.Logger
	client *redis.Client
}

var redisCon *redisConnecion

/*
   ConnectRedis :
*/
func ConnectRedis(logger *zap.Logger, config config.Redis) (Redis, error) {
	logger.Info("Connecting to redis")
	client := redis.NewClient(&redis.Options{
		Addr:     config.Addr,
		Password: config.Password,
		DB:       config.DB,
	})
	_, err := client.Ping().Result()
	if err != nil {
		return nil, err
	}

	redisCon = &redisConnecion{
		logger: logger,
		client: client,
	}

	return redisCon, nil
}

// Close : Disconnect from redis.
func (conn *redisConnecion) Close() {
	conn.logger.Info("Disconnecting from redis")
	conn.client.Close()
}

/*
   GetClient :
*/
func (conn *redisConnecion) GetClient() *redis.Client {
	return conn.client
}
