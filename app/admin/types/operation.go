package types

import (
	"../models"
	"github.com/graphql-go/graphql"
)

var OperationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Operation",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.ID,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				md, _ := p.Source.(models.Operation)
				return md.Id.Hex(), nil
			},
		},
		"name": &graphql.Field{
			Type: graphql.String,
		},
		"description": &graphql.Field{
			Type: graphql.String,
		},
		"crud": &graphql.Field{
			Type: graphql.Int,
		},
	},
})

var OperationInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "OperationInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"name": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"description": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"crud": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
	},
})
