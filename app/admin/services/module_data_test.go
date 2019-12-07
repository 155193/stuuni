package services_test

import (
	"../../db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var idMo = primitive.NewObjectID()

var CasesAddModule = []db.CaseTest{
	{
		Name: "Only add Module",
		Model: bson.M{
			"_id":   idMo,
			"name":  "test",
			"icon":  "test",
			"color": "test",
		},
		Thereiserror: false,
	},
	{
		Name: "Add with equal ID",
		Model: bson.M{
			"_id":   idMo,
			"name":  "error",
			"icon":  "error",
			"color": "error",
		},
		Thereiserror: true,
	},
}

var CasesUpdateModule = []db.CaseTest{
	{
		Name: "Update existed Module",
		Id:   idMo,
		Model: bson.M{
			"_id,omitempty": idMo,
			"name":          "update",
			"icon":          "update",
			"color":         "update",
		},
		Thereiserror: false,
	},
	{
		Name: "Update non existed Module",
		Id:   primitive.NewObjectID(),
		Model: bson.M{
			"name":  "error",
			"icon":  "error",
			"color": "error",
		},
		Thereiserror: true,
	},
}

var CasesGetModules = []db.CaseTest{
	{
		Name:         "Get modules",
		Thereiserror: false,
	},
}

var CasesModulesofUser = []db.CaseTest{
	{
		Name:         "Get modules of user",
		IdList:       append(roles, idRole),
		Thereiserror: false,
	},
}

var CasesRemoveModule = []db.CaseTest{
	{
		Name:         "Remove existed Module",
		Id:           idMo,
		Thereiserror: false,
	},
	{
		Name:         "Remove non existed Module",
		Id:           primitive.NewObjectID(),
		Thereiserror: true,
	},
}
