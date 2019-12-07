package services

import (
	"../../db"
	"../models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ColOperations = "operations"
)

func AddOperation(oper bson.M) error {
	c, errC := db.GetCollection(ColOperations)

	if errC != nil {
		return errC
	}
	_, errC = c.InsertOne(db.Ctx, oper)
	return errC
}

func UpdateOperation(id primitive.ObjectID, oper bson.M) error {
	c, errC := db.GetCollection(ColOperations)

	if errC != nil {
		return errC
	}
	_, errC = c.UpdateOne(db.Ctx, bson.M{"_id": id}, bson.M{"$set": oper}, )
	return errC
}

func GetOperations() (*[]models.Operation, error) {
	c, errC := db.GetCollection(ColOperations)

	if errC != nil {
		return nil, errC
	}
	var results []models.Operation
	cursor, err := c.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &results)
	return &results, err
}

func GetOperationByID(idOperation primitive.ObjectID) (*models.Operation, error) {
	c, errC := db.GetCollection(ColOperations)

	if errC != nil {
		return nil, errC
	}
	var result models.Operation
	var err = c.FindOne(db.Ctx, bson.M{"_id": idOperation}).Decode(&result)
	return &result, err
}

func GetOperationsByIDs(ids []primitive.ObjectID) (*[]models.Operation, error) {
	c, errC := db.GetCollection(ColOperations)

	if errC != nil {
		return nil, errC
	}
	var results []models.Operation
	cursor, err := c.Find(context.TODO(), bson.M{"_id": bson.M{"$in": ids}})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &results)
	return &results, err
}

func GetOperationByName(name string) (*models.Operation, error) {
	c, errC := db.GetCollection(ColOperations)

	if errC != nil {
		return nil, errC
	}
	var result models.Operation
	var err = c.FindOne(db.Ctx, bson.M{"name": name}).Decode(&result)
	return &result, err
}

func RemoveOperation(idOperation primitive.ObjectID) (bool, error) {
	c, errC := db.GetCollection(ColOperations)

	if errC != nil {
		return false, errC
	}
	_, err := c.DeleteOne(db.Ctx,
		bson.M{"_id": idOperation},
	)
	return err == nil, err
}
