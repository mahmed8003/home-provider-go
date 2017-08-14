package app

import (
	"home-provider/config"
	"home-provider/db"

	"github.com/go-redis/redis"
	"github.com/rs/zerolog"
)

/*
Context :
*/
type Context struct {
	env    string
	config config.AppConfig
	logger zerolog.Logger
	db     db.Database
	redis  *redis.Client
}

/*
NewContext :
*/
func NewContext(env string, config config.AppConfig, logger zerolog.Logger, db db.Database, redis *redis.Client) Context {
	return Context{
		env:    env,
		config: config,
		logger: logger,
		db:     db,
		redis:  redis,
	}
}

/*
Env :
*/
func (c Context) Env() string {
	return c.env
}

/*
Config :
*/
func (c Context) Config() config.AppConfig {
	return c.config
}

/*
Logger :
*/
func (c Context) Logger() zerolog.Logger {
	return c.logger
}

/*
Db :
*/
func (c Context) Db() db.Database {
	return c.db
}

/*
Redis :
*/
func (c Context) Redis() *redis.Client {
	return c.redis
}
