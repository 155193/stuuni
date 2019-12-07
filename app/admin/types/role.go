package types

import (
	"../models"
	"github.com/graphql-go/graphql"
)

var RoleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Role",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				md, _ := p.Source.(models.Role)
				return md.Id.Hex(), nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"accesses": &graphql.Field{
			Type: graphql.NewList(AccessType),
		},
	},
})

var RoleInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "RoleInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})
