package mutations

import (
	secServices "../../security/services"
	"../models"
	"../services"
	"../types"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var AddRoleMutation = &graphql.Field{
	Type:        types.RoleType, // the return type for this field
	Description: "Add View",
	Args: graphql.FieldConfigArgument{
		"idModule": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
		"role": &graphql.ArgumentConfig{
			Type: types.RoleInputType,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		if errPerm := secServices.ValidatePermissionToOperation(params); errPerm != nil {
			return nil, errPerm
		}
		module, role, roleBson, err := models.RoleAddToBson(params.Args["idModule"], params.Args["role"])
		defer secServices.AddOperationAudit(params, "", "", err)
		if err != nil {
			return nil, err
		}
		err = services.AddRole(module.Id, *roleBson)
		if err != nil {
			return nil, err
		}
		return *role, err
	},
}

var UpdateRoleMutation = &graphql.Field{
	Type:        types.RoleType,
	Description: "Update Role",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
		"role": &graphql.ArgumentConfig{
			Type: types.RoleInputType,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		if errPerm := secServices.ValidatePermissionToOperation(params); errPerm != nil {
			return nil, errPerm
		}
		role, err := models.RoleUpdateToBson(params.Args["id"], params.Args["role"])
		defer secServices.AddOperationAudit(params, "", "", err)
		if err != nil {
			return nil, err
		}
		err = services.UpdateRole(role.Id, role.Name, role.Description)
		if err != nil {
			return nil, err
		}
		return *role, err
	},
}

var RemoveRoleMutation = &graphql.Field{
	Type:        graphql.Boolean,
	Description: "Remove Role",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//verify permmissions of user to do operation
		if errPerm := secServices.ValidatePermissionToOperation(params); errPerm != nil {
			return nil, errPerm
		}
		id, _ := params.Args["id"].(string)
		idr, _ := primitive.ObjectIDFromHex(id)
		bol, err := services.RemoveRole(idr)
		//audit graphql action
		defer secServices.AddOperationAudit(params, "", "", err)
		return bol, err
	},
}
