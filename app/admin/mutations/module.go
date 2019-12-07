package mutations

import (
	secServices "../../security/services"
	"../models"
	"../services"
	"../types"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var AddModuleMutation = &graphql.Field{
	Type:        types.ModuleType, // the return type for this field
	Description: "Add Module",
	Args: graphql.FieldConfigArgument{
		"module": &graphql.ArgumentConfig{
			Type: types.ModuleInputType,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//verify permissions of user to do operation
		if errPerm := secServices.ValidatePermissionToOperation(params); errPerm != nil {
			return nil, errPerm
		}
		//convert module to Bson
		module, moduleBson, err := models.ModuleToBson(params.Args["module"])
		//audit graphql action
		defer secServices.AddOperationAudit(params, "", "", err)
		if err != nil {
			return nil, err
		}
		//add module service
		err = services.AddModule(*moduleBson)
		if err != nil {
			return nil, err
		}
		return *module, err
	},
}

var UpdateModuleMutation = &graphql.Field{
	Type:        types.ModuleType, // the return type for this field
	Description: "Update Module",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
		"module": &graphql.ArgumentConfig{
			Type: types.ModuleInputType,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//verify permissions of user to do operation
		if errPerm := secServices.ValidatePermissionToOperation(params); errPerm != nil {
			return nil, errPerm
		}
		//convert module to Bson
		module, moduleBson, err := models.ModuleUpdateToBson(params.Args["id"], params.Args["module"])
		//audit graphql action
		defer secServices.AddOperationAudit(params, "", "", err)
		if err != nil {
			return nil, err
		}
		//update module service
		err = services.UpdateModule(module.Id, *moduleBson)
		if err != nil {
			return nil, err
		}
		//get module updated
		result, err := services.GetModule(module.Id)
		if err != nil {
			return nil, err
		}
		return *result, err
	},
}

var RemoveModuleMutation = &graphql.Field{
	Type:        graphql.Boolean, // the return type for this field
	Description: "Remove Module",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//verify permissions of user to do operation
		if errPerm := secServices.ValidatePermissionToOperation(params); errPerm != nil {
			return nil, errPerm
		}
		id, _ := params.Args["id"].(string)
		idm, _ := primitive.ObjectIDFromHex(id)
		bol, err := services.RemoveModule(idm)
		//audit graphql action
		defer secServices.AddOperationAudit(params, "", "", err)
		return bol, err
	},
}
