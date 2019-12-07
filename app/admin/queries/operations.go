package queries

import (
	secServices "../../security/services"
	"../services"
	"../types"
	"github.com/graphql-go/graphql"
)

var OperationsQuery = &graphql.Field{
	Type:        graphql.NewList(types.OperationType),
	Description: "Get all Operations",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		if errPerm := secServices.ValidatePermissionToOperation(params); errPerm != nil {
			return nil, errPerm
		}
		operations, err := services.GetOperations()
		defer secServices.AddOperationAudit(params, "", "", err)
		if err != nil {
			return nil, err
		}
		return *operations, err
	},
}
