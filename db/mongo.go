package db

import (
	"home-provider/config"

	"github.com/rs/zerolog"

	mgo "gopkg.in/mgo.v2"
)

type dbConnecion struct {
	logger  zerolog.Logger
	session *mgo.Session
}

var dbCon *dbConnecion

/*
ConnectMongo :
*/
func ConnectMongo(logger zerolog.Logger, config config.Database) error {

	logger.Info().Msg("Connecting to database")
	session, err := mgo.Dial(config.Url)
	dbCon = &dbConnecion{
		logger:  logger,
		session: session,
	}
	return err
}

/*
GetDB :
*/
func GetDB() *mgo.Session {
	return dbCon.session
}

/*
DisconnectMongo :
*/
func DisconnectMongo() {
	if dbCon != nil {
		dbCon.logger.Info().Msg("Disconnecting from database")
		dbCon.session.Close()
	}
}
