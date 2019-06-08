package gql

import (
	"errors"
	mongo "valuebetsmining/src/src/mongodb"

	"github.com/graphql-go/graphql"
)

//Resolv ... Resolv the struct team
type Resolv struct {
	db *mongo.DriverMongo
}

var (
	//ErrParams ... Error parsing params
	ErrParams = errors.New("Error parsing params")
)

//TeamResolv ... Resolv a team
func (r *Resolv) TeamResolv(p graphql.ResolveParams) (interface{}, error) {

	local, ok := p.Args["local"].(string) //assertion
	if !ok {
		return nil, ErrParams
	}
	away, ok := p.Args["away"].(string) //assertion
	if !ok {
		return nil, ErrParams
	}
	cols, ok := p.Args["cols"].(string) //assertion
	if ok {
		users, err := r.db.GetProperties(local, away, cols)
		if err != nil {
			return nil, err
		}
		return users, nil
	}

	return nil, nil
}
