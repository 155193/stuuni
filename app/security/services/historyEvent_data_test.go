package services_test

import (
	"../../db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

var idHis = primitive.NewObjectID()

var CasesAddHistoryEvent = []db.CaseTest{
	{
		Name: "Only add Historyevent",
		Model: bson.M{
			"_id":  idHis,
			"name": "test",
			"date": time.Now().Format(time.RFC3339),
			"user": idHis,
			"info": "test",
		},
		Thereiserror: false,
	},
	{
		Name: "Add with equal id",
		Id:   idHis,
		Model: bson.M{
			"_id":  idHis,
			"name": "test",
			"date": time.Now().Format(time.RFC3339),
			"user": idHis,
			"info": "test",
		},
		Thereiserror: true,
	},
}

var CasesGetHistoryEvents = []db.CaseTest{
	{
		Name:         "Get History Events",
		Thereiserror: false,
	},
}
