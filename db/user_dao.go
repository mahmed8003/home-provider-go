package db

import (
	"home-provider/models"

	"gopkg.in/mgo.v2/bson"

	"gopkg.in/mgo.v2"
)

/*
mongoDB :
*/
type userDao struct {
	c *mgo.Collection
}

/*
NewUserDao :
*/
func NewUserDao(c *mgo.Collection) UserDao {

	return &userDao{
		c: c,
	}
}

// CreateUser : Create a new user in database
func (db *userDao) CreateUser(u *models.User) error {
	id := bson.NewObjectId()
	u.Id = id
	return db.c.Insert(u)
}
