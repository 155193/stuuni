package services_test

import (
	"../../db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var idRo = primitive.NewObjectID()
var roles []primitive.ObjectID

var CasesAddRole = []db.CaseTest{
	{
		Name: "Only add Role",
		Id:   idModule,
		Model: bson.M{
			"_id":         idRo,
			"name":        "test",
			"description": "test",
		},
		Thereiserror: false,
	},
	//{
	//	Name: "Add Role with existed id",
	//	Id:  idModule,
	//	Model: bson.M{
	//		"_id":         idRo,
	//		"name":        "test",
	//		"description": "test",
	//	},
	//	Thereiserror: true,
	//},
}

var CasesUpdateRole = []db.CaseTest{
	{
		Name: "Update existed Role",
		Id:   idRo,
		Nam:  "update",
		Des:  "update",
		Model: bson.M{
			"_id,omitempty": idRo,
			"name":          "update",
			"description":   "update",
			//"accesses":      nil,
		},
		Thereiserror: false,
	},
	{
		Name:         "Update non existed Role",
		Id:           primitive.NewObjectID(),
		Nam:          "error",
		Des:          "error",
		Thereiserror: true,
	},
}

var CasesGetRolesofUser = []db.CaseTest{
	{
		Name:         "Get Roles of user",
		IdList:       append(roles, idRo),
		Thereiserror: false,
	},
}

var CasesRemoveRole = []db.CaseTest{
	{
		Name:         "Remove exited Role",
		Id:           idRo,
		Thereiserror: false,
	},
	{
		Name:         "Remove non existed",
		Id:           primitive.NewObjectID(),
		Thereiserror: true,
	},
}
