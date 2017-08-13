package app

import (
	"home-provider/config"

	"github.com/go-redis/redis"
	"github.com/rs/zerolog"
	mgo "gopkg.in/mgo.v2"
)

/*
Context :
*/
type Context struct {
	env      string
	config   config.AppConfig
	logger   zerolog.Logger
	database *mgo.Session
	redis    *redis.Client
}

/*
NewContext :
*/
func NewContext(env string, config config.AppConfig, logger zerolog.Logger, database *mgo.Session, redis *redis.Client) Context {
	return Context{
		env:      env,
		config:   config,
		logger:   logger,
		database: database,
		redis:    redis,
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
Database :
*/
func (c Context) Database() *mgo.Session {
	return c.database
}

/*
Redis :
*/
func (c Context) Redis() *redis.Client {
	return c.redis
}
