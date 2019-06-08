package gql

import (
	mongo "valuebetsmining/src/src/mongodb"

	"github.com/graphql-go/graphql"
)

//Root ... Handling object of graphql
type Root struct {
	Query *graphql.Object
}

//NewRoot ... Return a fresh instance of Root struct
func NewRoot(db *mongo.DriverMongo) *Root {

	resolv := Resolv{db: db}

	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"users": &graphql.Field{
						Type: graphql.NewList(Properties),
						Args: graphql.FieldConfigArgument{
							"name": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
						},
						Resolve: resolv.TeamResolv,
					},
				},
			},
		),
	}

	return &root

}
