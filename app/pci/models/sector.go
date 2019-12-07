package models

import (
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"regexp"
)

type Sector struct {
	Id      primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	KmFrom  string             `bson:"kmFrom"`
	DesFrom string             `bson:"desFrom"`
	KmTo    string             `bson:"kmTo"`
	DesTo   string             `bson:"desTo"`
}

var R = regexp.MustCompile("[0-9][+][0-9]+")

func SectorAddToBson(sectorI interface{}) (primitive.ObjectID, *bson.M, error) {
	sectorM, ok := sectorI.(map[string]interface{})
	if !ok {
		return primitive.NewObjectID(), nil, errors.New("Error casting sector")
	}
	id := primitive.NewObjectID()
	sectorBson := bson.M{
		"_id":     id,
		"kmFrom":  sectorM["kmFrom"].(string),
		"desFrom": sectorM["desFrom"].(string),
		"kmTo":    sectorM["kmTo"].(string),
		"desTo":   sectorM["desTo"].(string),
	}
	return id, &sectorBson, nil
}
