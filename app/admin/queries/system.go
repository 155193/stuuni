package queries

import (
	"github.com/graphql-go/graphql"
	"time"
)

var GetDateTimeQuery = &graphql.Field{
	Type:        graphql.String,
	Description: "Get Date Time of server",
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		return time.Now().Format(time.RFC3339), nil
	},
}
