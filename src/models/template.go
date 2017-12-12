package models

import "gopkg.in/mgo.v2/bson"

type Template struct {
	ID   bson.ObjectId `bson:"_id" json:"id"`
	Name string        `bson:"name" json:"name"`
	Path string        `bson:"path" json:"path"`
}
