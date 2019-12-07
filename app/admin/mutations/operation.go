package mutations

import (
	secServices "../../security/services"
	"../models"
	"../services"
	"../types"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var AddOperationMutation = &graphql.Field{
	Type:        types.OperationType, // the return type for this field
	Description: "Add Operation",
	Args: graphql.FieldConfigArgument{
		"operation": &graphql.ArgumentConfig{
			Type: types.OperationInputType,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		if errPerm := secServices.ValidatePermissionToOperation(params); errPerm != nil {
			return nil, errPerm
		}
		oper, operBson, err := models.OperationAddToBson(params.Args["operation"])
		defer secServices.AddOperationAudit(params, "", "", err)
		if err != nil {
			return nil, err
		}
		err = services.AddOperation(*operBson)
		if err != nil {
			return nil, err
		}
		return *oper, err
	},
}

var UpdateOperationMutation = &graphql.Field{
	Type:        types.OperationType,
	Description: "Update Operation",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
		"operation": &graphql.ArgumentConfig{
			Type: types.OperationInputType,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		if errPerm := secServices.ValidatePermissionToOperation(params); errPerm != nil {
			return nil, errPerm
		}
		id, oper, err := models.OperationUpdateToBson(params.Args["id"], params.Args["operation"])
		defer secServices.AddOperationAudit(params, "", "", err)
		if err != nil {
			return nil, err
		}
		err = services.UpdateOperation(*id, *oper)
		if err != nil {
			return nil, err
		}
		operNew, err := services.GetOperationByID(*id)
		if err != nil {
			return nil, err
		}
		return *operNew, err
	},
}

var RemoveOperationMutation = &graphql.Field{
	Type:        graphql.Boolean,
	Description: "Remove Operation",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		if errPerm := secServices.ValidatePermissionToOperation(params); errPerm != nil {
			return nil, errPerm
		}
		id, _ := params.Args["id"].(string)
		ido, _ := primitive.ObjectIDFromHex(id)
		bol, err := services.RemoveOperation(ido)
		defer secServices.AddOperationAudit(params, "", "", err)
		return bol, err
	},
}
