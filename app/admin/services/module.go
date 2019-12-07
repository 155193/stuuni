package services

import (
	"../../db"
	"../models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ColModules = "modules"
)

func AddModule(module bson.M) error {
	c, errC := db.GetCollection(ColModules)
	if errC != nil {
		return errC
	}
	_, errC = c.InsertOne(db.Ctx, module)
	return errC
}

func UpdateModule(id primitive.ObjectID, module bson.M) error {
	c, errC := db.GetCollection(ColModules)

	if errC != nil {
		return errC
	}
	_, errC = c.UpdateOne(db.Ctx, bson.M{"_id": id}, bson.M{"$set": module}, )
	return errC
}

func GetModules() (*[]models.Module, error) {
	c, errC := db.GetCollection(ColModules)

	if errC != nil {
		return nil, errC
	}
	var results []models.Module
	cursor, err := c.Find(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	err = cursor.All(context.TODO(), &results)
	return &results, err
}

func GetModule(idModule primitive.ObjectID) (*models.Module, error) {
	c, errC := db.GetCollection(ColModules)

	if errC != nil {
		return nil, errC
	}
	var result models.Module
	var err = c.FindOne(db.Ctx, bson.M{"_id": idModule}).Decode(&result)
	return &result, err
}

func RemoveModule(idModule primitive.ObjectID) (bool, error) {
	c, errC := db.GetCollection(ColModules)

	if errC != nil {
		return false, errC
	}
	_, err := c.DeleteOne(db.Ctx,
		bson.M{"_id": idModule},
	)
	return err == nil, err
}

func ModulesOfUser(roles []primitive.ObjectID) (*[]models.Module, error) {
	pipeline := []bson.M{
		{
			"$match": bson.M{"roles._id": bson.M{"$in": roles}},
		},
		{
			"$project": bson.M{
				"_id":   1,
				"name":  1,
				"icon":  1,
				"color": 1,
				"roles": bson.M{
					"$filter": bson.M{
						"input": "$roles",
						"as":    "role",
						"cond": bson.M{
							"$in": []interface{}{"$$role._id", roles},
						},
					},
				},
			},
		},
	}
	c, errC := db.GetCollection(ColModules)

	if errC != nil {
		return nil, errC
	}
	var results []models.Module
	cursor, err := c.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return &results, nil
}

func ModulesOperationByUser(idOper primitive.ObjectID, roles []primitive.ObjectID) (*[]models.Module, error) {
	pipeline := []bson.M{
		{
			"$match": bson.M{
				"views.operations": idOper,
				"roles._id":        bson.M{"$in": roles},
			},
		},
	}
	c, errC := db.GetCollection(ColModules)

	if errC != nil {
		return nil, errC
	}
	var results []models.Module
	cursor, err := c.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return &results, nil
}
