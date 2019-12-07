package mutations

import (
	ms "../../security/models"
	secServices "../../security/services"
	"../models"
	"../services"
	"../types"
	"github.com/graphql-go/graphql"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var AddUserMutation = &graphql.Field{
	Type:        types.UserType, // the return type for this field
	Description: "Add User",
	Args: graphql.FieldConfigArgument{
		"user": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(types.UserInputType),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//verify permissions of user to do operation
		//if errPerm := secServices.ValidatePermissionToOperation(params); errPerm != nil {
		//	return nil, errPerm
		//}
		//convert user to Bson
		user, userBson, err := models.UserToBson(params.Args["user"])
		//audit graphql action
		//defer secServices.AddOperationAudit(params, "", "", err)
		if err != nil {
			return nil, err
		}
		//add user service
		err = services.AddUser(*userBson)
		if err != nil {
			return nil, err
		}
		return *user, err
	},
}

var UpdateUserMutation = &graphql.Field{
	Type:        types.UserType,
	Description: "Update User",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
		"user": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(types.UserInputType),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//verify permmissions of user to do operation
		//if errPerm := secServices.ValidatePermissionToOperation(params); errPerm != nil {
		//	return nil, errPerm
		//}
		//convert user to Bson
		user, userBson, err := models.UserUpdateToBson(params.Args["id"], params.Args["user"])
		//audit graphql action
		//defer secServices.AddOperationAudit(params, "", "", err)
		if err != nil {
			return nil, err
		}
		//update user service
		err = services.UpdateUser(user.Id, *userBson)
		if err != nil {
			return nil, err
		}
		//get user updated
		result, err := services.GetUserById(user.Id)
		if err != nil {
			return nil, err
		}
		return *result, err
	},
}

var RemoveUserMutation = &graphql.Field{
	Type:        graphql.Boolean,
	Description: "Remove User",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//verify permmissions of user to do operation
		//if errPerm := secServices.ValidatePermissionToOperation(params); errPerm != nil {
		//	return nil, errPerm
		//}
		id, _ := params.Args["id"].(string)
		idu, _ := primitive.ObjectIDFromHex(id)
		bol, err := services.RemoveUser(idu)
		//audit graphql action
		//defer secServices.AddOperationAudit(params, "", "", err)
		return bol, err
	},
}

var AddRole2UserMutation = &graphql.Field{
	Type:        types.UserType,
	Description: "Add Role to User",
	Args: graphql.FieldConfigArgument{
		"idUser": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
		"idRole": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//verify permmissions of user to do operation
		//if errPerm := secServices.ValidatePermissionToOperation(params); errPerm != nil {
		//	return nil, errPerm
		//}
		idUser, _ := params.Args["idUser"].(string)
		idu, _ := primitive.ObjectIDFromHex(idUser)
		idRole, _ := params.Args["idRole"].(string)
		idr, _ := primitive.ObjectIDFromHex(idRole)
		bol, err := services.AddRole2User(idu, idr)
		//audit graphql action
		//defer secServices.AddOperationAudit(params, "", "", err)
		if bol == false {
			return nil, err
		}
		result, err := services.GetUserById(idu)
		if err != nil {
			return nil, err
		}
		return *result, err
	},
}

var RemoveRole2UserMutation = &graphql.Field{
	Type:        types.UserType,
	Description: "Remove Role Of User",
	Args: graphql.FieldConfigArgument{
		"idUser": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
		"idRole": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//verify permmissions of user to do operation
		//if errPerm := secServices.ValidatePermissionToOperation(params); errPerm != nil {
		//	return nil, errPerm
		//}
		idUser, _ := params.Args["idUser"].(string)
		idu, _ := primitive.ObjectIDFromHex(idUser)
		idRole, _ := params.Args["idRole"].(string)
		idr, _ := primitive.ObjectIDFromHex(idRole)
		bol, err := services.RemoveRole2User(idu, idr)
		//audit graphql action
		//defer secServices.AddOperationAudit(params, "", "", err)
		if bol == false {
			return nil, err
		}
		result, err := services.GetUserById(idu)
		if err != nil {
			return nil, err
		}
		return *result, err
	},
}

var RefreshPasswordUserMutation = &graphql.Field{
	Type:        types.UserType,
	Description: "Refresh Password to default dni",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.ID),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//verify permissions of user to do operation
		//if errPerm := secServices.ValidatePermissionToOperation(params); errPerm != nil {
		//	return nil, errPerm
		//}
		id, _ := params.Args["id"].(string)
		idu, _ := primitive.ObjectIDFromHex(id)
		user, err := services.RefreshPasswordUser(idu)
		//audit graphql action
		defer secServices.AddOperationAudit(params, "", "", err)
		if err != nil {
			return nil, err
		}
		return *user, nil
	},
}

var UpdatePasswordUserMutation = &graphql.Field{
	Type:        types.UserType,
	Description: "Update Password of User",
	Args: graphql.FieldConfigArgument{
		"password": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"newPassword": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//verify permissions of user to do operation
		//if errPerm := secServices.ValidatePermissionToOperation(params); errPerm != nil {
		//	return nil, errPerm
		//}
		user := params.Context.Value("current").(ms.Values).Get("currentUser").(interface{})
		id := user.(services.UserTkn).Id
		idu, _ := primitive.ObjectIDFromHex(id)
		password, _ := params.Args["password"].(string)
		newPassword, _ := params.Args["newPassword"].(string)
		u, err := services.UpdatePasswordUser(idu, password, newPassword)
		//audit graphql action
		//defer secServices.AddOperationAudit(params, "", "", err)
		if err != nil {
			return nil, err
		}
		return *u, err
	},
}

var AddUserGenericMutation = &graphql.Field{
	Type:        types.UserType, // the return type for this field
	Description: "Add User",
	Args: graphql.FieldConfigArgument{
		"userGeneric": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(types.UserGenericInputType),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//verify permissions of user to do operation
		//if errPerm := secServices.ValidatePermissionToOperation(params); errPerm != nil {
		//	return nil, errPerm
		//}
		//convert user to Bson
		id, userBson, err := models.UserGenericToBson(params.Args["userGeneric"])
		//audit graphql action
		//defer secServices.AddOperationAudit(params, "", "", err)
		if err != nil {
			return nil, err
		}
		//add user service
		err = services.AddUser(*userBson)
		if err != nil {
			return nil, err
		}
		userGeneric, err := services.GetUserById(*id)
		if err != nil {
			return nil, err
		}
		return *userGeneric, nil
	},
}
