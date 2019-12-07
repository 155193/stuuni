package types

import (
	adminServices "../../admin/services"
	adminTypes "../../admin/types"
	"../models"
	"github.com/graphql-go/graphql"
)

var AuditEventType = graphql.NewObject(graphql.ObjectConfig{
	Name: "AuditEvent",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				auditEvent, _ := p.Source.(models.AuditEvent)
				return auditEvent.Id.Hex(), nil
			},
		},
		"user": &graphql.Field{
			Type: adminTypes.UserType,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				auditEvent, _ := p.Source.(models.AuditEvent)
				user,err:=adminServices.GetUserById(auditEvent.User)
				if err!=nil {
					return nil,nil
				}
				return *user,nil
			},
		},
		"ipLocal": &graphql.Field{
			Type: graphql.String,
		},
		"ipRemote": &graphql.Field{
			Type: graphql.String,
		},
		"agent": &graphql.Field{
			Type: graphql.String,
		},
		"auth": &graphql.Field{
			Type: graphql.String,
		},
		"token": &graphql.Field{
			Type: graphql.String,
		},
		"consult": &graphql.Field{
			Type: graphql.String,
		},
		"operations": &graphql.Field{
			Type: graphql.NewList(OperationAuditType),
		},
	},
})
