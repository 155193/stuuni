package types

import (
	"../models"
	"../services"
	"github.com/graphql-go/graphql"
)

var ViewType = graphql.NewObject(graphql.ObjectConfig{
	Name: "View",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				md, _ := p.Source.(models.View)
				return md.Id.Hex(), nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"url": &graphql.Field{
			Type: graphql.String,
		},
		"icon": &graphql.Field{
			Type: graphql.String,
		},
		"operations": &graphql.Field{
			Type: graphql.NewList(OperationType),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				md, _ := p.Source.(models.View)
				return services.GetOperationsByIDs(md.Operations)
			},
		},
	},
})

var ViewInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "ViewInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"url": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"icon": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		//"operations": &graphql.InputObjectFieldConfig{
		//	Type: graphql.NewNonNull(graphql.ID),
		//},
	},
})
