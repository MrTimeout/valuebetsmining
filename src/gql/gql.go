package gql

import (
	"../postgres"
	"github.com/graphql-go/graphql"
)

type Root struct {
	Query *graphql.Object
}

func NewRoot(db *postgres.Db) *Root {

	resolv := Resolv{db: db}

	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name:   "Query",
				Fields: graphql.Fields{
					"users": &graphql.Field{
						
					}
				},
			},
		),
	}

}