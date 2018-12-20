package graphql_util

import (
	"github.com/graphql-go/graphql"
	"github.com/OdaDaisuke/go-graphql-sample/graphql_util/fields"
)

var rootMutation = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootMutation",
	Fields: graphql.Fields{
		"createUser": fields.CreateUserField,
	},
})