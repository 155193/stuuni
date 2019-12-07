package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type HistoryEvent struct {
	Id   primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name string             `bson:"name"`
	Date string             `bson:"date"`
	User primitive.ObjectID `bson:"user"`
	Info string             `bson:"info"`
}
