package types

import (
	"../models"
	"../services"
	"github.com/graphql-go/graphql"
)

var AccessType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Access",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				md, _ := p.Source.(models.Access)
				return md.Id.Hex(), nil
			},
		},
		"view": &graphql.Field{
			Type: ViewType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				md, _ := p.Source.(models.Access)
				v, err := services.ViewOfUser(md.View)
				return *v, err
			},
		},
		"order": &graphql.Field{
			Type: graphql.Int,
		},
		"position": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var AccessInputType = graphql.NewInputObject(graphql.InputObjectConfig{
	Name: "AccessInput",
	Fields: graphql.InputObjectConfigFieldMap{
		"view": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
		"order": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.Int),
		},
		"position": &graphql.InputObjectFieldConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
})
