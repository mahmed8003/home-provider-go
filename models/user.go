package models

import "gopkg.in/mgo.v2/bson"

// User holds metadata about a user.
type User struct {
	Id bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
}
