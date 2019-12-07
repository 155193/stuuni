package services

import (
	"../../db"
	"../models"
	"context"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var (
	ColRoles = "modules"
)

func AddRole(idModule primitive.ObjectID, data2set bson.M) error {
	c, errC := db.GetCollection(ColRoles)

	if errC != nil {
		return errC
	}
	_, errC = c.UpdateOne(db.Ctx,
		bson.M{"_id": idModule},
		bson.M{
			"$addToSet": bson.M{
				"roles": data2set,
			},
		},
	)
	return errC
}

func UpdateRole(idRole primitive.ObjectID, name string, description string) error {
	c, errC := db.GetCollection(ColRoles)

	if errC != nil {
		return errC
	}
	_, errC = c.UpdateOne(db.Ctx,
		bson.M{"roles._id": idRole},
		bson.M{"$set": bson.M{
			"roles.$.name":        name,
			"roles.$.description": description,
		}},
	)
	return errC
}

func GetRole(idRole primitive.ObjectID) (*models.Role, error) {
	c, errC := db.GetCollection(ColAccess)

	if errC != nil {
		return nil, errC
	}
	var result models.Role
	pipeline := []bson.M{
		{
			"$unwind": "$roles",
		},
		{"$match": bson.M{"roles._id": idRole}},
		{
			"$project": bson.M{
				"_id":         "$roles._id",
				"name":        "$roles.name",
				"description": "$roles.description",
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

func RemoveRole(idRole primitive.ObjectID) (bool, error) {
	c, errC := db.GetCollection(ColRoles)

	if errC != nil {
		return false, errC
	}
	_, err := c.UpdateOne(db.Ctx,
		bson.M{"roles._id": idRole},
		bson.M{"$pull": bson.M{
			"roles": bson.M{
				"_id": idRole,
			},
		}},
	)
	return err == nil, err
}

func GetRolesOfUser(roles []primitive.ObjectID) (*[]models.Role, error) {
	pipeline := []bson.M{
		{
			"$match": bson.M{"roles._id": bson.M{"$in": roles}},
		},
		{
			"$project": bson.M{
				"_id": 0,
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
		{
			"$unwind": "$roles",
		},
		{
			"$project": bson.M{
				"_id":         "$roles._id",
				"name":        "$roles.name",
				"description": "$roles.description",
				"accesses":    "$roles.accesses",
			},
		},
	}

	c, errC := db.GetCollection(ColRoles)

	if errC != nil {
		return nil, errC
	}
	var results []models.Role
	cursor, err := c.Aggregate(context.TODO(), pipeline)
	if err != nil {
		return nil, err
	}
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}
	return &results, nil
}

// get index of role
func GetIndexOfRole(idRole primitive.ObjectID) (int, error) {
	c, errC := db.GetCollection(ColModules)

	if errC != nil {
		return -1, errC
	}
	var module models.Module
	errC = c.FindOne(db.Ctx, bson.M{"roles._id": idRole}).Decode(&module)
	if errC != nil {
		return -1, errC
	}
	if len(module.Roles) == 0 {
		return -1, errors.New("error not found role")
	}
	for iRole, role := range module.Roles {
		if role.Id == idRole {
			return iRole, nil
		}
	}
	return -1, errors.New("error not found role")
}
