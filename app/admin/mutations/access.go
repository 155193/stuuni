package mutations

import (
	secServices "../../security/services"
	"../models"
	"../services"
	"../types"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var AddAccessMutation = &graphql.Field{
	Type:        types.AccessType,
	Description: "Add Access",
	Args: graphql.FieldConfigArgument{
		"idModule": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
		"idRole": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
		"access": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(types.AccessInputType),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		if errPerm := secServices.ValidatePermissionToOperation(params); errPerm != nil {
			return nil, errPerm
		}
		module, role, access, accessBson, err := models.AccessAddToBson(params.Args["idModule"],
			params.Args["idRole"], params.Args["access"])
		defer secServices.AddOperationAudit(params, "", "", err)
		if err != nil {
			return nil, err
		}
		if err := services.AddAccess(module.Id, role.Id, *accessBson); err != nil {
			return nil, err
		}
		return *access, nil
	},
}

var UpdateAccessMutation = &graphql.Field{
	Type:        types.AccessType,
	Description: "Update Access",
	Args: graphql.FieldConfigArgument{
		"idRole": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
		"idAccess": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
		"access": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(types.AccessInputType),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		if errPerm := secServices.ValidatePermissionToOperation(params); errPerm != nil {
			return nil, errPerm
		}
		idR, idA, idV, order, position, err := models.AccessUpdateToBson(params.Args["idRole"], params.Args["idAccess"], params.Args["access"])
		defer secServices.AddOperationAudit(params, "", "", err)
		if err != nil {
			return nil, err
		}
		err = services.UpdateAccess(idR, idA, idV, order, position)
		if err != nil {
			return nil, err
		}
		access, err := services.GetAccess(idA)
		if err != nil {
			return nil, err
		}
		return *access, nil
	},
}

var RemoveAccessMutation = &graphql.Field{
	Type:        graphql.Boolean,
	Description: "Remove Access",
	Args: graphql.FieldConfigArgument{
		"idRole": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
		"idAccess": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		if errPerm := secServices.ValidatePermissionToOperation(params); errPerm != nil {
			return nil, errPerm
		}
		idRole, _ := params.Args["idRole"].(string)
		idr, _ := primitive.ObjectIDFromHex(idRole)
		idAccess, _ := params.Args["idAccess"].(string)
		ida, _ := primitive.ObjectIDFromHex(idAccess)
		bol, err := services.RemoveAccess(idr, ida)
		defer secServices.AddOperationAudit(params, "", "", err)
		return bol, err
	},
}
