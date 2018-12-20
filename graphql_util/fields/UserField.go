package fields

import (
	"errors"
	"github.com/graphql-go/graphql"
	"github.com/OdaDaisuke/go-graphql-sample/graphql_util/types"
	"github.com/OdaDaisuke/go-graphql-sample/services"
	"github.com/OdaDaisuke/go-graphql-sample/models"
)

var UserField = &graphql.Field{
	Type: types.UserType,
	Description: "Get single user",
	Args: graphql.FieldConfigArgument{
		"id": &graphql.ArgumentConfig{
			Type: graphql.String,
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		userId, isOK := params.Args["id"].(string)
		if isOK {
			return services.FindByUserId(userId)
		}
		return nil, errors.New("no userId")
	},
}

var CreateUserField = &graphql.Field{
	Type: types.UserType,
	Description: "Create new user",
	Args: graphql.FieldConfigArgument{
		"userId": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
		"email": &graphql.ArgumentConfig{
			Type: graphql.NewNonNull(graphql.String),
		},
	},
	Resolve: func(params graphql.ResolveParams) (interface{}, error) {
		userName, _ := params.Args["userName"].(string)
		email, _ := params.Args["email"].(string)

		newUser, err := models.CreateUser(userName, email)
		if err != nil {
			panic(err)
		}
		return newUser, nil
	},
}