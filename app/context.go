package app

import (
	"home-provider/config"
	"home-provider/db"

	"go.uber.org/zap"
)

/*
Context :
*/
type Context interface {
	Env() string
	Config() *config.AppConfig
	Logger() *zap.Logger
	Redis() db.Redis
	Db() db.Database
}

/*
appContext :
*/
type appContext struct {
	env    string
	config *config.AppConfig
	logger *zap.Logger
	redis  db.Redis
	db     db.Database
}

/*
NewContext :
*/
func NewContext(env string, config *config.AppConfig, logger *zap.Logger, redis db.Redis, db db.Database) Context {
	return &appContext{
		env:    env,
		config: config,
		logger: logger,
		redis:  redis,
		db:     db,
	}
}

func (ctx *appContext) Env() string {
	return ctx.env
}

func (ctx *appContext) Config() *config.AppConfig {
	return ctx.config
}

func (ctx *appContext) Logger() *zap.Logger {
	return ctx.logger
}

func (ctx *appContext) Redis() db.Redis {
	return ctx.redis
}

func (ctx *appContext) Db() db.Database {
	return ctx.db
}
