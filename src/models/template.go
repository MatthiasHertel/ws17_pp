package models

import "gopkg.in/mgo.v2/bson"

type Template struct {
	ID    bson.ObjectId `bson:"_id" json:"id"`
	JobID string        `bson:"jobid" json:"jobid"`
	Name  string        `bson:"name" json:"name"`
	Path  string        `bson:"path" json:"path"`
}
