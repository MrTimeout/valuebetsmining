package models

import (
	"context"
	"strings"
	"time"
	mongo "valuebetsmining/src/mongodb"
	"valuebetsmining/src/mongodb/entities"

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
func GetPropertiesOfATeam(codiv, team string) (entities.Team, error) {
	driver, err := IsDBCOl(codiv)
	if err != nil {
		return entities.Team{}, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.D{
		primitive.E{Key: "Date", Value: "test"},
		primitive.E{Key: "LocalTeam", Value: team},
	}
	options := options.FindOptions{}
	limit := int64(1)
	options.Limit = &limit
	options.Sort = bson.D{primitive.E{Key: "Index", Value: -1}}
	collections, err := driver.Client.Database(mongo.DBDbase).Collection(codiv).Find(ctx, filter, &options)
	if err != nil {
		return entities.Team{}, err
	}
	defer collections.Close(ctx)
	result := entities.Result{}
	for collections.Next(ctx) {
		err := collections.Decode(&result)
		if err != nil {
			return entities.Team{}, err
		}
	}
	return entities.NewTeam(result), nil
}

//GetPropertiesOfALocalTeam ... Get all properties of a local team
func GetPropertiesOfALocalTeam(codiv, team string) (entities.TeamLocal, error) {
	driver, err := IsDBCOl(codiv)
	if err != nil {
		return entities.TeamLocal{}, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.D{
		primitive.E{Key: "Date", Value: "test"},
		primitive.E{Key: "LocalTeam", Value: team},
	}
	options := options.FindOptions{}
	limit := int64(1)
	options.Limit = &limit
	options.Sort = bson.D{primitive.E{Key: "Index", Value: -1}}
	collections, err := driver.Client.Database(mongo.DBDbase).Collection(codiv).Find(ctx, filter, &options)
	if err != nil {
		return entities.TeamLocal{}, err
	}
	defer collections.Close(ctx)
	result := entities.Result{}
	for collections.Next(ctx) {
		err := collections.Decode(&result)
		if err != nil {
			return entities.TeamLocal{}, err
		}
	}
	return entities.NewLocalTeam(result), nil
}

//GetPropertiesOfAAwayTeam ... Get all properties of a Away team
func GetPropertiesOfAAwayTeam(codiv, team string) (entities.TeamAway, error) {
	driver, err := IsDBCOl(codiv)
	if err != nil {
		return entities.TeamAway{}, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	filter := bson.D{
		primitive.E{Key: "Date", Value: "test"},
		primitive.E{Key: "AwayTeam", Value: team},
	}
	options := options.FindOptions{}
	limit := int64(1)
	options.Limit = &limit
	options.Sort = bson.D{primitive.E{Key: "Index", Value: -1}}
	collections, err := driver.Client.Database(mongo.DBDbase).Collection(codiv).Find(ctx, filter, &options)
	if err != nil {
		return entities.TeamAway{}, err
	}
	defer collections.Close(ctx)
	result := entities.Result{}
	for collections.Next(ctx) {
		err := collections.Decode(&result)
		if err != nil {
			return entities.TeamAway{}, err
		}
	}
	return entities.NewAwayTeam(result), nil
}
