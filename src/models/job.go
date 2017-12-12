package models

import "gopkg.in/mgo.v2/bson"

// Represents a Job, bson for naming the mongo properties
type Job struct {
	ID          bson.ObjectId `bson:"_id" json:"id"`
	Name        string        `bson:"name" json:"name"`
	Description string        `bson:"description" json:"description"`
	Param       string        `bson:"param" json:"param"`
	Templates   []Template    `bson:"templates" json:"templates"`
}
