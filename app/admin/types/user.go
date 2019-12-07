package types

import (
	"../models"
	"../services"
	"github.com/graphql-go/graphql"
)

var UserType = graphql.NewObject(graphql.ObjectConfig{
	Name: "User",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				md, _ := p.Source.(models.User)
				return md.Id.Hex(), nil
			},
		},
		"names": &graphql.Field{
			Type: graphql.String,
		},
		"ln": &graphql.Field{
			Type: graphql.String,
		},
		"mln": &graphql.Field{
			Type: graphql.String,
		},
		"dni": &graphql.Field{
			Type: graphql.String,
		},
		"photo": &graphql.Field{
			Type: graphql.String,
		},
		"nick": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
		//"password": &graphql.Field{
		//	Type: graphql.String,
		//},
		"token": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				md, _ := p.Source.(models.User)
				token, _ := services.CreateTokenString(md.Id.Hex())
				if token == "" {
					return nil, nil
				}
				return token, nil
			},
		},
		"modules": &graphql.Field{
			Type: graphql.NewList(ModuleType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				md, _ := p.Source.(models.User)
				modules, err := services.ModulesOfUser(md.Roles)
				if err != nil {
					return nil, err
				}
				return *modules, nil
			},
		},
		"roles": &graphql.Field{
			Type: graphql.NewList(RoleType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				user, _ := p.Source.(models.User)
				return services.GetRolesOfUser(user.Roles)
			},
		},
	},
})

var UserInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UserInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"names": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"ln": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"mln": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"dni": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"photo": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"nick": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"email": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		//"password": &graphql.InputObjectFieldConfig{
		//	Type: graphql.String,
		//},
	},
})
// User Generic
var UserGenericInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "UserGenericInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"names": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"ln": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"mln": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"dni": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})
