package resolvers

import (
	mongo "valuebetsmining/src/mongodb"

	"github.com/graphql-go/graphql"
)

//Resolv ... Struct that handles mongo database
type Resolv struct {
	db *mongo.DriverMongo
}

//ErrorResolv ... Struct that handles error of struct resolv
type ErrorResolv struct {
	ErrorString string
}

var (
	//ErrInvalidCo ... Error parsing country because of the format
	ErrInvalidCo = &ErrorResolv{ErrorString: "Error parsing country because dont follow the good format"}
	//ErrInvalidDiv ... Error parsing division because of the format
	ErrInvalidDiv = &ErrorResolv{ErrorString: "Error parsing division because dont follow the good format"}
	//ErrNil ... Error parsing nil argumet
	ErrNil = &ErrorResolv{ErrorString: "Error parsing nil argument"}
)

//TeamPropertiesResolv ... Method that handles properties of each team
func (r *Resolv) TeamPropertiesResolv(p graphql.ResolveParams) (interface{}, error) {
	country, ok := p.Args["country"].(string)
	if !ok {
		return nil, ErrNil
	}
	division, ok := p.Args["division"].(string)
	if !ok {
		return nil, ErrNil
	}
	team, ok := p.Args["team"].(string)
	if !ok {
		return nil, ErrNil
	}
	where, ok := p.Args["type"].(string) //all, away, local
	if !ok {
		where = "all"
	}
	return resul, nil
}

//Error ... Return error of struct Resolv
func (er *ErrorResolv) Error() string {
	return er.ErrorString
}
