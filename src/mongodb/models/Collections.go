package models

import (
	"context"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
	mongo "valuebetsmining/src/mongodb"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

var (
	//ErrNotFoundData ... Error not found data
	ErrNotFoundData = errors.New("Data didnt found in the database")
)

//GetAllCollectionNames ... Get all collection names of an existing database
func GetAllCollectionNames(db string) ([]string, error) {
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
	res, err := driver.Client.Database(mongo.DBDbase).ListCollections(ctx, bsonx.Doc{})
	if err != nil {
		return nil, err
	}
	str := make([]string, 0)
	for res.Next(ctx) {
		next := &bsonx.Doc{}
		err = res.Decode(next)
		if err != nil {
			return nil, err
		}

		elem, err := next.LookupErr("name")
		if err != nil {
			return nil, err
		}

		if elem.Type() != bson.TypeString {
			return nil, fmt.Errorf("incorrect type for 'name'. got %v. want %v", elem.Type(), bson.TypeString)
		}

		elemName := elem.StringValue()
		str = append(str, elemName)
	}
	return str, nil
}

//CountryDiv ... Return a map of country:[]divisions
func CountryDiv() (map[string][]string, error) {
	cols, err := GetAllCollectionNames(mongo.DBDbase)
	if err != nil {
		return nil, err
	}
	re := regexp.MustCompile(`[A-Z]{1,2}([a-z]{1,}|[0-9]?)`)
	finalMap := make(map[string][]string)
	for _, val := range cols {
		temp := re.FindAllString(val, -1)
		if len(temp) >= 2 {
			finalMap[temp[0]] = temp[1:]
		}
	}
	return finalMap, nil
}

//Countries ... Return countries in the database
func Countries() ([]string, error) {
	m, err := CountryDiv()
	if err != nil {
		return nil, err
	}
	resl := make([]string, 0, len(m))
	for k := range m {
		resl = append(resl, k)
	}
	return resl, nil
}

//Divisions ... Return all divisions of a country
func Divisions(c string) ([]string, error) {
	m, err := CountryDiv()
	if err != nil {
		return nil, err
	}
	resl := make([]string, 0, len(m))
	for _, v := range m[c] {
		resl = append(resl, v)
	}
	if len(resl) == 0 {
		return nil, ErrNotFoundData
	}
	return resl, nil
}

//AmIHereCol ... Return true if the col exists and false otherwise
func AmIHereCol(col string) (bool, error) {
	if strings.Trim(col, " ") == "" || len(strings.Trim(col, " ")) == 0 {
		return false, mongo.ErrEmptyString
	}
	driver, err := mongo.ConnectDB()
	if err != nil {
		return false, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	str := []string{}
	res, err := driver.Client.Database(mongo.DBDbase).ListCollections(ctx, bson.D{primitive.E{Key: "name", Value: col}})
	if err != nil {
		return false, err
	}
	for res.Next(ctx) {
		next := &bsonx.Doc{}
		err = res.Decode(next)
		if err != nil {
			return false, err
		}

		elem, err := next.LookupErr("name")
		if err != nil {
			return false, err
		}

		if elem.Type() != bson.TypeString {
			return false, fmt.Errorf("incorrect type for 'name'. got %v. want %v", elem.Type(), bson.TypeString)
		}

		elemName := elem.StringValue()
		str = append(str, elemName)
	}
	return len(str) == 1, nil
}

//AmIHereDB ... Return true if the db exists and false otherwise
func AmIHereDB(db string) (bool, error) {
	if strings.Trim(db, " ") == "" || len(strings.Trim(db, " ")) == 0 {
		return false, mongo.ErrEmptyString
	}
	driver, err := mongo.ConnectDB()
	if err != nil {
		return false, err
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	dbs, err := driver.Client.ListDatabaseNames(ctx, bson.D{primitive.E{Key: "name", Value: db}})
	if err != nil {
		return false, err
	}
	return len(dbs) == 1, nil
}
