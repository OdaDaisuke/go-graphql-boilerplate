package graphql_util

import (
	"github.com/graphql-go/graphql"
	"github.com/OdaDaisuke/go-graphql-sample/graphql_util/fields"
)

var rootQuery = graphql.NewObject(graphql.ObjectConfig{
	Name: "RootQuery",
	Fields: graphql.Fields{ // Modelに該当
		"user": fields.UserField,
	},
})