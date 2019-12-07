package services

import (
	"../../db"
	"../models"
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	Colview = "modules"
)

func AddView(idModule primitive.ObjectID, view bson.M) error {
	c, errC := db.GetCollection(Colview)

	if errC != nil {
		return errC
	}
	_, errC = c.UpdateOne(db.Ctx,
		bson.M{"_id": idModule},
		bson.M{
			"$addToSet": bson.M{
				"views": view,
			},
		},
	)
	return errC
}

func UpdateView(idView primitive.ObjectID, nam string, des string, url string, icon string) error {
	c, errC := db.GetCollection(Colview)

	if errC != nil {
		return errC
	}
	_, errC = c.UpdateOne(db.Ctx,
		bson.M{"views._id": idView},
		bson.M{"$set": bson.M{
			"views.$.name":        nam,
			"views.$.description": des,
			"views.$.url":         url,
			"views.$.icon":        icon,
		}},
	)
	return errC
}

func GetView(idView primitive.ObjectID) (*models.View, error) {
	c, errC := db.GetCollection(Colview)

	if errC != nil {
		return nil, errC
	}
	var result models.View
	pipeline := []bson.M{
		{
			"$unwind": "$views",
		},
		{"$match": bson.M{"views._id": idView}},
		{
			"$project": bson.M{
				"_id":         "$views._id",
				"name":        "$views.name",
				"description": "$views.description",
				"url":         "$views.url",
				"icon":        "$views.icon",
			},
		},
	}
	cursor, err := c.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	if err = cursor.Decode(&result); err != nil {
		return nil, err
	}
	return &result, nil
}

func RemoveView(idView primitive.ObjectID) (bool, error) {
	c, errC := db.GetCollection(Colview)

	if errC != nil {
		return false, errC
	}
	_, err := c.UpdateOne(db.Ctx,
		bson.M{"views._id": idView},
		bson.M{"$pull": bson.M{
			"views": bson.M{
				"_id": idView,
			},
		}},
	)
	return err == nil, err
}

func GetViewinAccess(id primitive.ObjectID) (bool, error) {
	c, errC := db.GetCollection(Colview)

	if errC != nil {
		return false, errC
	}
	var result models.Module
	err := c.FindOne(db.Ctx, bson.M{
		"roles.accesses.view": id,
	}).Decode(&result)
	return err == nil, err
}

func ViewOfUser(idView primitive.ObjectID) (*models.View, error) {
	pipeline := []bson.M{
		{
			"$match": bson.M{"views._id": idView},
		},
		{
			"$project": bson.M{
				"views": bson.M{
					"$filter": bson.M{
						"input": "$views",
						"as":    "view",
						"cond": bson.M{
							"$eq": []interface{}{"$$view._id", idView},
						},
					},
				},
			},
		},
		{
			"$unwind": "$views",
		},
		{
			"$project": bson.M{
				"_id":         "$views._id",
				"name":        "$views.name",
				"description": "$views.description",
				"url":         "$views.url",
				"icon":        "$views.icon",
				"operations":  "$views.operations",
			},
		},
	}
	c, errC := db.GetCollection(Colview)

	if errC != nil {
		return nil, errC
	}
	var view models.View
	cursor, err := c.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	if err = cursor.Decode(&view); err != nil {
		return nil, err
	}
	return &view, nil
}

func AddOperation2View(idView primitive.ObjectID, idOperation primitive.ObjectID) error {
	c, errC := db.GetCollection(ColAccess)

	if errC != nil {
		return errC
	}
	_, errC = c.UpdateOne(db.Ctx,
		bson.M{"views._id": idView},
		bson.M{
			"$addToSet": bson.M{
				"views.$.operations": idOperation,
			},
		},
	)
	return errC
}

func RemoveOperation2View(idView primitive.ObjectID, idOperation primitive.ObjectID) (bool, error) {
	c, errC := db.GetCollection(ColAccess)

	if errC != nil {
		return false, errC
	}
	_, err := c.UpdateOne(db.Ctx,
		bson.M{"views._id": idView},
		bson.M{"$pull": bson.M{
			"views.$.operations": idOperation,
		}},
	)
	return err == nil, err
}
