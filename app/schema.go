package main

import (
	adminMutations "./admin/mutations"
	adminQueries "./admin/queries"
	secQueries "./security/queries"
	"github.com/graphql-go/graphql"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{
		// admin
		"users":       adminQueries.UsersQuery,
		"user":        adminQueries.UserQuery,
		"modules":     adminQueries.ModulesQuery,
		"operations":  adminQueries.OperationsQuery,
		"logint":      adminQueries.LoginTQuery,
		"getDateTime": adminQueries.GetDateTimeQuery,
		// security
		"auditEvents": secQueries.AuditEventsQuery,
		// projects

	},
})

// root mutation
var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		// admin
		"addModule":             adminMutations.AddModuleMutation,
		"updateModule":          adminMutations.UpdateModuleMutation,
		"removeModule":          adminMutations.RemoveModuleMutation,
		"addView":               adminMutations.AddViewMutation,
		"updateView":            adminMutations.UpdateViewMutation,
		"removeView":            adminMutations.RemoveViewMutation,
		"addRole":               adminMutations.AddRoleMutation,
		"updateRole":            adminMutations.UpdateRoleMutation,
		"removeRole":            adminMutations.RemoveRoleMutation,
		"addAccess":             adminMutations.AddAccessMutation,
		"updateAccess":          adminMutations.UpdateAccessMutation,
		"removeAccess":          adminMutations.RemoveAccessMutation,
		"addUser":               adminMutations.AddUserMutation,
		"addUserGeneric":        adminMutations.AddUserGenericMutation,
		"updateUser":            adminMutations.UpdateUserMutation,
		"removeUser":            adminMutations.RemoveUserMutation,
		"addRole2User":          adminMutations.AddRole2UserMutation,
		"removeRole2User":       adminMutations.RemoveRole2UserMutation,
		"refreshPasswordUser":   adminMutations.RefreshPasswordUserMutation,
		"updatePasswordUser":    adminMutations.UpdatePasswordUserMutation,
		"addOperation":          adminMutations.AddOperationMutation,
		"updateOperation":       adminMutations.UpdateOperationMutation,
		"removeOperation":       adminMutations.RemoveOperationMutation,
		"addOperation2View":     adminMutations.AddOperation2ViewMutation,
		"removeOperationOfView": adminMutations.RemoveOperationOfViewMutation,
		// security
		//"addHistoryEvent": secMutations.AddHistoryEventMutation,
	},
})

func GetSchema() *graphql.Schema {

	var schema, _ = graphql.NewSchema(graphql.SchemaConfig{
		Query:    rootQuery,
		Mutation: rootMutation,
	})
	return &schema
}
