package services_test

import (
	"../../db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var idUse = primitive.NewObjectID()

var CasesAddUser = []db.CaseTest{
	{
		Name: "Only add user",
		Model: bson.M{
			"_id":      idUse,
			"names":    "test",
			"ln":       "test",
			"mln":      "test",
			"dni":      "12345678",
			"photo":    "test",
			"nick":     "test",
			"email":    "test@test.com",
			"password": "test",
		},
		Thereiserror: false,
	},
	{
		Name: "Add with equal Id",
		Model: bson.M{
			"_id":      idUse,
			"names":    "test",
			"ln":       "test",
			"mln":      "test",
			"dni":      "12345678",
			"photo":    "test",
			"nick":     "test",
			"email":    "test@test.com",
			"password": "test",
		},
		Thereiserror: true,
	},
}

var CasesUpdateUser = []db.CaseTest{
	{
		Name: "Update existed User",
		Id:   idUse,
		Model: bson.M{
			"_id,omitempty": idUse,
			"names":         "update",
			"ln":            "update",
			"mln":           "update",
			"dni":           "12345678",
			"photo":         "update",
			"nick":          "update",
			"email":         "update@update.com",
			"password":      "update",
		},
		Thereiserror: false,
	},
	{
		Name: "Update non existed User",
		Id:   primitive.NewObjectID(),
		Model: bson.M{
			"_id": idUse,
		},
		Thereiserror: true,
	},
}

var CasesAddRole2User = []db.CaseTest{
	{
		Name:         "Add role to existed user",
		Id:           idUse,
		Idaux:        idRole,
		Thereiserror: false,
	},
	{
		Name:         "Add role to non existed user",
		Id:           primitive.NewObjectID(),
		Idaux:        primitive.NewObjectID(),
		Thereiserror: true,
	},
}

var CasesRemoveRole2User = []db.CaseTest{
	{
		Name:         "Remove role to existed user",
		Id:           idUse,
		Idaux:        idRole,
		Thereiserror: false,
	},
	{
		Name:         "Remove role to non existed user",
		Id:           primitive.NewObjectID(),
		Idaux:        idRole,
		Thereiserror: true,
	},
	//{
	//	Name:         "Remove non existed Role to user",
	//	Id:           idUse,
	//	Idaux:        primitive.NewObjectID(),
	//	Thereiserror: true,
	//},
}

var CasesGetuserbyId = []db.CaseTest{
	{
		Name:         "Get existed user",
		Id:           idUse,
		Thereiserror: false,
	},
	{
		Name:         "Get non existed User",
		Id:           primitive.NewObjectID(),
		Thereiserror: true,
	},
}

var CasesGetUserbyDni = []db.CaseTest{
	{
		Name:         "Get existed user",
		Dni:          "12345678",
		Thereiserror: false,
	},
	{
		Name:         "Get non existed user",
		Dni:          "74185296",
		Thereiserror: true,
	},
}

var CasesGetUsers = []db.CaseTest{
	{
		Name:         "Get users",
		Thereiserror: false,
	},
}

var CasesLogin = []db.CaseTest{
	{
		Name:         "Login with existed account",
		Dni:          "update",
		Thereiserror: false,
	},
	{
		Name:         "Login with non existed account",
		Dni:          "sigetest",
		Thereiserror: true,
	},
}

var CasesLoginT = []db.CaseTest{
	{
		Name:         "Login with existed id",
		Dni:          idUse.Hex(),
		Thereiserror: false,
	},
	{
		Name:         "Login with non existed id",
		Dni:          primitive.NewObjectID().Hex(),
		Thereiserror: true,
	},
}

var CaseCreateTokenString = []db.CaseTest{
	{
		Name:         "Create token",
		Dni:          idUse.String(),
		Thereiserror: false,
	},
}

var CasesDecodeTokenString = []db.CaseTest{
	{
		Name:         "Decode existed token",
		Dni:          token,
		Thereiserror: false,
	},
	{
		Name:         "Decode non existed token",
		Dni:          "テスト",
		Thereiserror: true,
	},
}

var CasesUpdatePasswordUser = []db.CaseTest{
	{
		Name:         "Update password of existed user",
		Id:           idUse,
		Dni:          "12345678",
		Icon:         "updateagain",
		Thereiserror: false,
	},
	{
		Name:         "Update of non existed user",
		Id:           primitive.NewObjectID(),
		Dni:          "update",
		Icon:         "プルエバ:(",
		Thereiserror: true,
	},
}

var CasesRefreshPasswordUser = []db.CaseTest{
	{
		Name:         "Refresh passowrd of existed user",
		Id:           idUse,
		Thereiserror: false,
	},
	{
		Name:         "Refresh password of non existed user",
		Id:           primitive.NewObjectID(),
		Thereiserror: true,
	},
}

var CasesRemoveUser = []db.CaseTest{
	{
		Name:         "Remove existed user",
		Id:           idUse,
		Thereiserror: false,
	},
	{
		Name:         "Remove non existed user",
		Id:           primitive.NewObjectID(),
		Thereiserror: true,
	},
}
