package services_test

import (
	"../../db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var idView = primitive.NewObjectID()

var CasesAddView = []db.CaseTest{
	{
		Name: "Only Add",
		Id:   idModule,
		Model: bson.M{
			"_id":         idView,
			"name":        "test",
			"description": "test",
			"url":         "test",
			"icon":        "test",
		},
		Thereiserror: false,
	},
	//{
	//	Name: "Add with equal ID",
	//	Id:  idModule,
	//	Model: bson.M{
	//		"_id":         idView,
	//		"name":        "error",
	//		"description": "error",
	//		"url":         "error",
	//		"icon":        "error",
	//	},
	//	Thereiserror: true,
	//},
}

var CasesUpdateView = []db.CaseTest{
	{
		Name: "Update existed View",
		Id:   idView,
		Nam:  "update",
		Des:  "update",
		Url:  "update",
		Icon: "update",
		Model: bson.M{
			"_id,omitempty": idView,
			"name":          "update",
			"description":   "update",
			"url":           "update",
			"icon":          "update",
		},
		Thereiserror: false,
	},
	{
		Name:         "Update non existed View",
		Id:           primitive.NewObjectID(),
		Nam:          "error",
		Des:          "error",
		Url:          "error",
		Icon:         "error",
		Thereiserror: true,
	},
}

var CasesRemoveView = []db.CaseTest{
	{
		Name:         "Remove existed View",
		Id:           idView,
		Thereiserror: false,
	},
	{
		Name:         "Remove non existed View",
		Id:           primitive.NewObjectID(),
		Thereiserror: true,
	},
}

var CasesViewofUser = []db.CaseTest{
	{
		Name:         "Get exited view of user",
		Id:           idView,
		Thereiserror: false,
	},
	{
		Name:         "Get non exited view of user",
		Id:           primitive.NewObjectID(),
		Thereiserror: true,
	},
}
