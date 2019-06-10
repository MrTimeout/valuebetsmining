package searcher

import (
	"valuebetsmining/src/gpl/resolvers"
	"valuebetsmining/src/gpl/types"
	mongo "valuebetsmining/src/mongodb"

	"github.com/graphql-go/graphql"
)

//Root .. Object of grahql to search by params
type Root struct {
	Query *graphql.Object
}

//NewRoot ... Return a fresh instance of an object root
func NewRoot(db *mongo.DriverMongo) *Root {

	resolv := resolvers.Resolv{db: "pepe"}

	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "PropTeam",
				Fields: graphql.Fields{
					"Properties": &graphql.Field{
						Type: graphql.NewList(types.Properties),
						Args: graphql.FieldConfigArgument{
							"country": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
							"division": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
							"team": &graphql.ArgumentConfig{
								Type: graphql.String,
							},
						},
						Resolve: db.Client,
					},
				},
			},
		),
	}

	return &root

}
