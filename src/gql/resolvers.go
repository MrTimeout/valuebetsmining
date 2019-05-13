package gql

import (
	"valuebetsmining/src/postgres"
	"github.com/graphql-go/graphql"
)

type Resolv struct {
	db *postgres.Db
}

func (r *Resolv) UserResolv(p graphql.ResolveParams) (interface{}, error) {

	name, ok := p.Args["name"].(string) //assertion
	if ok {
		users := r.db.GetUsersByName(name)
		return users, nil
	}

	return nil, nil
}