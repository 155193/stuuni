package mutations

import (
	secServices "../../security/services"
	"../models"
	"../services"
	"../types"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var AddViewMutation = &graphql.Field{
	Type:        types.ViewType, // the return type for this field
	Description: "Add View",
	Args: graphql.FieldConfigArgument{
		"idModule": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
		"view": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(types.ViewInputType),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//verify permissions of user to do operation
		if errPerm := secServices.ValidatePermissionToOperation(params); errPerm != nil {
			return nil, errPerm
		}
		//convert view to Bson
		module, view, viewBson, err := models.ViewToBson(params.Args["idModule"], params.Args["view"])
		//audit graphql action
		defer secServices.AddOperationAudit(params, "", "", err)
		if err != nil {
			return nil, err
		}
		//add view service
		err = services.AddView(module.Id, *viewBson)
		if err != nil {
			return nil, err
		}
		return *view, nil
	},
}

var UpdateViewMutation = &graphql.Field{
	Type:        types.ViewType,
	Description: "Update View",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
		"view": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(types.ViewInputType),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//verify permissions of user to do operation
		if errPerm := secServices.ValidatePermissionToOperation(params); errPerm != nil {
			return nil, errPerm
		}
		//convert view to Bson
		view, err := models.ViewUpdateToBson(params.Args["id"], params.Args["view"])
		//audit graphql action
		defer secServices.AddOperationAudit(params, "", "", err)
		if err != nil {
			return nil, err
		}
		//update view service
		err = services.UpdateView(view.Id, view.Name, view.Description, view.Url, view.Icon)
		if err != nil {
			return nil, err
		}
		return *view, err
	},
}

var RemoveViewMutation = &graphql.Field{
	Type:        graphql.Boolean,
	Description: "Remove View",
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
		idv, _ := primitive.ObjectIDFromHex(id)
		//get view in access service
		bol, err := services.GetViewinAccess(idv)
		//audit graphql action
		defer secServices.AddOperationAudit(params, "", "", err)
		if bol {
			return false, err
		}
		//remove view service
		bol, err = services.RemoveView(idv)
		return bol, err
	},
}

var AddOperation2ViewMutation = &graphql.Field{
	Type:        graphql.Boolean,
	Description: "Add Operation to View",
	Args: graphql.FieldConfigArgument{
		"idView": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
		"idOperation": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//verify permissions of user to do operation
		if errPerm := secServices.ValidatePermissionToOperation(params); errPerm != nil {
			return nil, errPerm
		}
		idView, _ := params.Args["idView"].(string)
		idv, _ := primitive.ObjectIDFromHex(idView)
		idOperation, _ := params.Args["idOperation"].(string)
		ido, _ := primitive.ObjectIDFromHex(idOperation)
		err := services.AddOperation2View(idv, ido)
		//audit graphql action
		defer secServices.AddOperationAudit(params, "", "", err)
		if err != nil {
			return nil, err
		}
		return true, nil
	},
}

var RemoveOperationOfViewMutation = &graphql.Field{
	Type:        graphql.Boolean,
	Description: "Remove Operation of view",
	Args: graphql.FieldConfigArgument{
		"idView": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
		"idOperation": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//verify permissions of user to do operation
		if errPerm := secServices.ValidatePermissionToOperation(params); errPerm != nil {
			return nil, errPerm
		}
		idView, _ := params.Args["idView"].(string)
		idv, _ := primitive.ObjectIDFromHex(idView)
		idOperation, _ := params.Args["idOperation"].(string)
		ido, _ := primitive.ObjectIDFromHex(idOperation)
		bol, err := services.RemoveOperation2View(idv, ido)
		//audit graphql action
		defer secServices.AddOperationAudit(params, "", "", err)
		return bol, err
	},
}
