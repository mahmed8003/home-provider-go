package db

import (
	"home-provider/config"

	"go.uber.org/zap"

	mgo "gopkg.in/mgo.v2"
)

/*
mongoDB :
*/
type mongoDB struct {
	logger *zap.Logger
	conn   *mgo.Session
	db     *mgo.Database
	//
	userDao UserDao
}

/*
ConnectMongo :
*/
func ConnectMongo(logger *zap.Logger, config config.Database) (Database, error) {

	logger.Info("Connecting to database")
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

	//
	userColl := db.C("users")
	userDao := NewUserDao(userColl)
	//

	mongo := &mongoDB{
		logger:  logger,
		conn:    conn,
		db:      db,
		userDao: userDao,
	}

	logger.Info("Database connection successfull")
	return mongo, nil
}

// Close : Disconnect from database.
func (db *mongoDB) Close() {
	db.logger.Info("Disconnecting from database")
	db.conn.Close()
}

// GetUserDao : Returns.
func (db *mongoDB) GetUserDao() UserDao {
	return db.userDao
}
