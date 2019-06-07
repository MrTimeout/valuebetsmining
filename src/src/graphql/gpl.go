package gql

import (
	"github.com/graphql-go/graphql"
)

//Root ... Handling object of graphql
type Root struct {
	Query *graphql.Object
}

//NewRoot ... Return a fresh instance of Root struct
func NewRoot(db *postgres.Db) *Root {

	resolv := Resolv{db: db}

	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"users": &graphql.Field{
						Type: graphql.NewList(User),
						Args: graphql.FieldConfigArgument{
							"name": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
						},
						Resolve: resolv.UserResolv,
					},
				},
			},
		),
	}

	return &root

}
