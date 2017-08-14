package models

import "gopkg.in/mgo.v2/bson"

// User holds metadata about a user.
type User struct {
	Id       bson.ObjectId `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string        `json:"name,omitempty" bson:"name,omitempty"`
	Email    string        `json:"email,omitempty" bson:"email,omitempty"`
	Mobile   string        `json:"mobile,omitempty" bson:"mobile,omitempty"`
	Password string        `json:"password,omitempty" bson:"password,omitempty"`
}
