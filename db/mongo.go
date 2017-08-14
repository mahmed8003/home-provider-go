package db

import (
	"home-provider/config"

	"github.com/rs/zerolog"

	mgo "gopkg.in/mgo.v2"
)

/*
Database :
*/
type mongoDB struct {
	logger zerolog.Logger
	conn   *mgo.Session
	db     *mgo.Database
}

/*
ConnectMongo :
*/
func ConnectMongo(logger zerolog.Logger, config config.Database) (Database, error) {

	logger.Info().Msg("Connecting to database")
	conn, err := mgo.Dial(config.Uri)
	if err != nil {
		return nil, err
	}

	if config.Username != "" {
		cred := &mgo.Credential{
			Username: config.Username,
			Password: config.Password,
		}

		if err := conn.Login(cred); err != nil {
			return nil, err
		}
	}

	db := conn.DB(config.Database)
	mongo := &mongoDB{
		logger: logger,
		conn:   conn,
		db:     db,
	}

	logger.Info().Msg("Database connection successfull")
	return mongo, nil
}

// Close : Disconnect from database.
func (db *mongoDB) Close() {
	db.logger.Info().Msg("Disconnecting from database")
	db.conn.Close()
}
