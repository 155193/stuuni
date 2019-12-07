package queries

import (
	secServices "../../security/services"
	"../services"
	"../types"
	"github.com/graphql-go/graphql"
)

var ModulesQuery = &graphql.Field{
	Type:        graphql.NewList(types.ModuleType),
	Description: "Get all Modules",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//verify permissions of user to do operation
		if errPerm := secServices.ValidatePermissionToOperation(params); errPerm != nil {
			return nil, errPerm
		}
		//get modules service
		modules, err := services.GetModules()
		//audit graphql action
		defer secServices.AddOperationAudit(params, "", "", err)
		if err != nil {
			return nil, err
		}
		return *modules, nil
	},
}
