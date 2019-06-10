package models

import (
	"context"
	"strings"
	"time"
	mongo "valuebetsmining/src/src/mongodb"
	"valuebetsmining/src/src/mongodb/entities"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//GetAllTeamName ... Get all team names of an existing collection
func GetAllTeamName(codiv string) ([]interface{}, error) {
	if strings.Trim(mongo.DBDbase, " ") == "" || len(strings.Trim(mongo.DBDbase, " ")) == 0 {
		return nil, mongo.ErrEmptyString
	}
	driver, err := mongo.ConnectDB()
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	t, err := AmIHereDB(mongo.DBDbase)
	if err != nil {
		return nil, err
	} else if !t {
		return nil, mongo.ErrNotExistDB
	}
	t, err = AmIHereCol(codiv)
	if err != nil {
		return nil, err
	} else if !t {
		return nil, mongo.ErrNotExistCOL
	}
	filter := bson.D{
		primitive.E{Key: "Date", Value: primitive.Regex{Pattern: `^[0-9]{2}\/[0-9]{2}\/(18|19)$`, Options: ""}},
	}
	collections, err := driver.Client.Database(mongo.DBDbase).Collection(codiv).Distinct(ctx, "LocalTeam", filter)
	if err != nil {
		return nil, err
	}
	return collections, nil
}

//GetPropertiesOfATeam ... Get all properties of a team
func GetPropertiesOfATeam(codiv, team string) (entities.Result, error) {
	if strings.Trim(mongo.DBDbase, " ") == "" || len(strings.Trim(mongo.DBDbase, " ")) == 0 {
		return entities.Result{}, mongo.ErrEmptyString
	}
	driver, err := mongo.ConnectDB()
	if err != nil {
		return entities.Result{}, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	t, err := AmIHereDB(mongo.DBDbase)
	if err != nil {
		return entities.Result{}, err
	} else if !t {
		return entities.Result{}, mongo.ErrNotExistDB
	}
	t, err = AmIHereCol(codiv)
	if err != nil {
		return entities.Result{}, err
	} else if !t {
		return entities.Result{}, mongo.ErrNotExistCOL
	}
	filter := bson.D{
		primitive.E{Key: "Date", Value: primitive.Regex{Pattern: `^[0-9]{2}\/[0-9]{2}\/(18|19)$`, Options: ""}},
		primitive.E{Key: "LocalTeam", Value: team},
	}
	options := options.FindOptions{}
	limit := int64(1)
	options.Limit = &limit
	options.Sort = bson.D{primitive.E{Key: "Index", Value: -1}}
	collections, err := driver.Client.Database(mongo.DBDbase).Collection(codiv).Find(ctx, filter, &options)
	if err != nil {
		return entities.Result{}, err
	}
	defer collections.Close(ctx)
	result := entities.Result{}
	for collections.Next(ctx) {
		err := collections.Decode(&result)
		if err != nil {
			return entities.Result{}, err
		}
	}
	return result, nil
}
