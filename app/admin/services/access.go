package services

import (
	"../../db"
	"../models"
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ColAccess = "modules"
)

func AddAccess(idModule primitive.ObjectID, idRole primitive.ObjectID, access bson.M) error {
	c, errC := db.GetCollection(ColAccess)

	if errC != nil {
		return errC
	}
	_, errC = c.UpdateOne(db.Ctx,
		bson.M{"_id": idModule, "roles._id": idRole},
		bson.M{
			"$addToSet": bson.M{
				"roles.$.accesses": access,
			},
		},
	)
	return errC
}

func UpdateAccess(idRole primitive.ObjectID, idAccess primitive.ObjectID, idV primitive.ObjectID, order int, position string) error {
	c, errC := db.GetCollection(ColAccess)

	if errC != nil {
		return errC
	}
	// get index of role
	iRole, err := GetIndexOfRole(idRole)
	if err != nil {
		return err
	}
	_, errC = c.UpdateOne(db.Ctx,
		bson.M{"roles.accesses._id": idAccess},
		bson.M{"$set": bson.M{
			fmt.Sprintf("roles.%d.accesses.$.view", iRole):     idV,
			fmt.Sprintf("roles.%d.accesses.$.order", iRole):    order,
			fmt.Sprintf("roles.%d.accesses.$.position", iRole): position,
		}},
	)
	return errC
}

func RemoveAccess(idRole primitive.ObjectID, idAccess primitive.ObjectID) (bool, error) {
	c, errC := db.GetCollection(ColAccess)

	if errC != nil {
		return false, errC
	}
	_, err := c.UpdateOne(db.Ctx,
		bson.M{"roles._id": idRole},
		bson.M{"$pull": bson.M{
			"roles.$.accesses": bson.M{
				"_id": idAccess,
			},
		}},
	)
	return err == nil, err
}

func GetAccess(idAccess primitive.ObjectID) (*models.Access, error) {
	c, errC := db.GetCollection(ColAccess)

	if errC != nil {
		return nil, errC
	}
	var result models.Access
	pipeline := []bson.M{
		{
			"$unwind": "$roles",
		},
		{
			"$unwind": "$roles.accesses",
		},
		{
			"$match": bson.M{"roles.accesses._id": idAccess}},
		{
			"$project": bson.M{
				"_id":      "$roles.accesses._id",
				"view":     "$roles.accesses.view",
				"order":    "$roles.accesses.order",
				"position": "$roles.accesses.position",
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
