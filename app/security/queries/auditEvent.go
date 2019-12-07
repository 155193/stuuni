package queries

import (
	"../services"
	"../types"
	"github.com/graphql-go/graphql"
)

var AuditEventsQuery = &graphql.Field{
	Type:        graphql.NewList(types.AuditEventType),
	Description: "get audit events",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		audits, err := services.GetAuditEvents()
		if err != nil {
			return nil, err
		}
		return *audits, nil
	},
}
