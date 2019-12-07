package types

import (
	adminServices "../../admin/services"
	adminTypes "../../admin/types"
	"../models"
	"github.com/graphql-go/graphql"
)

var OperationAuditType = graphql.NewObject(graphql.ObjectConfig{
	Name: "OperationAudit",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				operationAudit, _ := p.Source.(models.OperationAudit)
				return operationAudit.Id.Hex(), nil
			},
		},
		"operation": &graphql.Field{
			Type: adminTypes.OperationType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				operationAudit, _ := p.Source.(models.OperationAudit)
				return adminServices.GetOperationByID(operationAudit.Operation)
			},
		},
		"permission": &graphql.Field{
			Type: graphql.Boolean,
		},
		"nowStruct": &graphql.Field{
			Type: graphql.String,
		},
		"newStruct": &graphql.Field{
			Type: graphql.String,
		},
		"error": &graphql.Field{
			Type: graphql.String,
		},
		"result": &graphql.Field{
			Type: graphql.Boolean,
		},
		"timeB": &graphql.Field{
			Type: graphql.String,
		},
		"timeE": &graphql.Field{
			Type: graphql.String,
		},
	},
})
