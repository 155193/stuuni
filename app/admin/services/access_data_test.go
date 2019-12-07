package services_test

import (
	"../../db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//DATA TEST CASES ACCESS

var idAccess = primitive.NewObjectID()

//Cases for addAccess
var casesAddAccess = []db.CaseTest{
	{
		Name:  "Only Add",
		Id:    idModule,
		Idaux: idRole,
		Model: bson.M{
			"_id":      idAccess,
			"view":     idAccess,
			"order":    1,
			"position": "test",
		},
		Thereiserror: false,
	},
	//{
	//	Name: "Add equal Id",
	//	Id:  idModule,
	//	Idaux:  idRole,
	//	Model: bson.M{
	//		"_id":      idAccess,
	//		"view":     idAccess,
	//		"order":    1,
	//		"position": "test",
	//	},
	//	Thereiserror: true,
	//},
}

//cases for updateAccess
var casesUpdateAccess = []db.CaseTest{
	{
		Name:     "Update existed Accesss",
		Id:       idAccess,
		Idaux:    idRole,
		Order:    2,
		Position: "update",
		Model: bson.M{
			//"_id,omitempty": idAccess,
			"view":          idRole,
			"order":         2,
			"position":      "update",
		},
		Thereiserror: false,
	},
	{
		Name:         "Update non existed Access",
		Id:           primitive.NewObjectID(),
		Idaux:        primitive.NewObjectID(),
		Order:        1,
		Position:     "-",
		Thereiserror: true,
	},
}

//cases for removeAccess
var casesRemoveAccess = []db.CaseTest{
	{
		Name:         "Remove existed access",
		Id:           idRole,
		Idaux:        idAccess,
		Thereiserror: false,
	},
	{
		Name:         "Remove non existed access",
		Id:           primitive.NewObjectID(),
		Idaux:        idAccess,
		Thereiserror: true,
	},
}
