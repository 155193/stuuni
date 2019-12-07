package services

import (
	"../../db"
	"../models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var Colhistory = "historyEvents"

func AddHistoryEvent(historyEvent bson.M) error {
	c, errC := db.GetCollection(Colhistory)
	if errC != nil {
		return errC
	}
	_, errC = c.InsertOne(db.Ctx, historyEvent)
	return errC
}

func GetHistoryEvents() (*[]models.HistoryEvent, error) {
	c, errC := db.GetCollection(Colhistory)

	if errC != nil {
		return nil, errC
	}
	var results []models.HistoryEvent
	cursor, err := c.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &results)
	return &results, err
}

func GetHistoryEvent(id primitive.ObjectID) (*models.HistoryEvent, error) {
	c, errC := db.GetCollection(Colhistory)

	if errC != nil {
		return nil, errC
	}
	var result models.HistoryEvent
	errC = c.FindOne(db.Ctx, bson.M{"_id": id}).Decode(&result)
	return &result, errC
}

func RemoveHistoryEvent(idHistoryEvent primitive.ObjectID) error {
	c, errC := db.GetCollection(Colhistory)

	if errC != nil {
		return errC
	}
	_, err := c.DeleteOne(db.Ctx,
		bson.M{"_id": idHistoryEvent},
	)
	return err
}
