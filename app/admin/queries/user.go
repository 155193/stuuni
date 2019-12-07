package queries

import (
	ms "../../security/models"
	"../services"
	"../types"
	"github.com/graphql-go/graphql"
)

var UsersQuery = &graphql.Field{
	Type:        graphql.NewList(types.UserType),
	Description: "Get All Users",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//if errPerm := secServices.ValidatePermissionToOperation(params); errPerm != nil {
		//	return nil, errPerm
		//}
		users, err := services.GetUsers()
		//defer secServices.AddOperationAudit(params, "", "", err)
		if err != nil {
			return nil, err
		}
		return *users, nil
	},
}

var LoginTQuery = &graphql.Field{
	Type:        types.UserType,
	Description: "Login User by nick and password",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//if errPerm := secServices.ValidatePermissionToOperation(params); errPerm != nil {
		//	return nil, errPerm
		//}
		user := params.Context.Value("current").(ms.Values).Get("currentUser").(interface{})
		use, err := services.LoginT(user.(services.UserTkn).Id)
		//defer secServices.AddOperationAudit(params, "", "", err)
		if err != nil {
			return nil, err
		}
		return *use, nil
	},
}

var UserQuery = &graphql.Field{
	Type:        types.UserType,
	Description: "Get user",
	Args: graphql.FieldConfigArgument{
		"dni": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		//if errPerm := secServices.ValidatePermissionToOperation(params); errPerm != nil {
		//	return nil, errPerm
		//}
		dni, _ := params.Args["dni"].(string)
		user, err := services.GetUserByDNI(dni)
		//defer secServices.AddOperationAudit(params, "", "", err)
		if err != nil {
			return nil, err
		}
		return *user, nil
	},
}
