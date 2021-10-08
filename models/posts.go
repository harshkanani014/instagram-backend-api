package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type Posts struct {
	Id        bson.ObjectId `json:"id" bson:"_id"`
	User      bson.ObjectId `json:"user" bson:"podcast,omitempty"`
	Caption   string        `json:"caption" bson:"caption"`
	ImageURL  string        `json:"imageURL" bson:"imageURL"`
	Timestamp time.Time     `json:"timestamp" bson:"timestamp"`
}
