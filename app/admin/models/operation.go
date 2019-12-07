package models

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Operation struct {
	Id          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Name        string             `bson:"name"`
	Description string             `bson:"description"`
	CRUD        int                `bson:"crud"`
}

func OperationAddToBson(operI interface{}) (*Operation, *bson.M, error) {
	operM, ok := operI.(map[string]interface{})
	if !ok {
		return nil, nil, errors.New("error casting")
	}
	id := primitive.NewObjectID()
	name := operM["name"].(string)
	description := operM["description"].(string)
	crud := operM["crud"].(int)
	operBson := bson.M{
		"_id":         id,
		"name":        name,
		"description": description,
		"crud":        crud,
	}
	return &Operation{Id: id, Name: name, CRUD: crud}, &operBson, nil
}

func OperationUpdateToBson(idV interface{}, operI interface{}) (*primitive.ObjectID, *bson.M, error) {
	ids, ok := idV.(string)
	if !ok {
		return nil, nil, errors.New("error casting Id operation")
	}
	id, err := primitive.ObjectIDFromHex(ids)
	if err != nil {
		return nil, nil, errors.New("error Id operation")
	}
	operM, ok := operI.(map[string]interface{})
	if !ok {
		return nil, nil, errors.New("error casting operation")
	}
	operBson := bson.M{
		"name":        operM["name"].(string),
		"description": operM["description"].(string),
		"crud":        operM["crud"].(int),
	}
	return &id, &operBson, nil
}
