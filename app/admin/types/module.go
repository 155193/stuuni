package types

import (
	"../models"
	"github.com/graphql-go/graphql"
)

var ModuleType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Module",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				md, _ := p.Source.(models.Module)
				return md.Id.Hex(), nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"icon": &graphql.Field{
			Type: graphql.String,
		},
		"color": &graphql.Field{
			Type: graphql.String,
		},
		"roles": &graphql.Field{
			Type: graphql.NewList(RoleType),
		},
		"views": &graphql.Field{
			Type: graphql.NewList(ViewType),
		},
	},
})

var ModuleInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "ModuleInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"icon": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"color": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})
