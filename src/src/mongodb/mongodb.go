package mongo

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

var (
	//DBHost ... Address of the database to connect to it
	DBHost = os.Getenv("MONGODB_ROOT_ADDR")
	//DBPort ... Port to connect to the database
	DBPort = os.Getenv("MONGO_INITDB_PORT")
	//DBUser ... User to connect to the database
	DBUser = os.Getenv("MONGO_INITDB_ROOT_USERNAME")
	//DBPassword ... Password to connect to the database
	DBPassword = os.Getenv("MONGO_INITDB_ROOT_PASSWORD")
	//DBDbase ... Database name to connect
	DBDbase = os.Getenv("MONGO_DB_NAME")
)

var (
	//ErrEmptyString ... Error parsing empty string
	ErrEmptyString = errors.New("Error parsing empty string")
	//ErrNotExistDB ... Error getting database, doesnt exist
	ErrNotExistDB = errors.New("Error getting database, doesnt exists")
	//ErrNotExistCOL ... Error parsing collection name
	ErrNotExistCOL = errors.New("Error parsing collection name")
)

//DriverMongo ... Driver mongodb
type DriverMongo struct {
	Client *mongo.Client
}

//Result ... Struct that handles information of the last match of each team
type Result struct {
	Last10WinningLocalMatchs,
	Last10TiedingLocalMatchs,
	Last10LosingLocalMatchs,
	Last10WinningAwayMatchs,
	Last10TiedingAwayMatchs,
	Last10LosingAwayMatchs,
	Last10StreackWinningLocal,
	Last10StreackNoLosingLocal,
	Last10StreackWinningAway,
	Last10StreackNoLosingAway int
	Last10AverageTuckedGoalsLocal,
	Last10AverageReceivedGoalsLocal,
	Last10AverageTuckedGoalsAway,
	Last10AverageReceivedGoalsAway float64
}

//ConnectDB ... Returns a fresh instance of a driver to connect to mongodb
func ConnectDB() (DriverMongo, error) {

	c, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(GetStringToConnect()))
	if err != nil {
		return DriverMongo{}, err
	}
	err = c.Ping(context.TODO(), nil)
	if err != nil {
		return DriverMongo{}, err
	}
	return DriverMongo{Client: c}, nil
}

//GetAllTeamName ... Get all team names of an existing collection
func (d *DriverMongo) GetAllTeamName(codiv string) ([]string, error) {
	if strings.Trim(DBDbase, " ") == "" || len(strings.Trim(DBDbase, " ")) == 0 {
		return nil, ErrEmptyString
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()
	t, err := d.AmIHereDB(DBDbase)
	if err != nil {
		return nil, err
	} else if !t {
		return nil, ErrNotExistDB
	}
	t, err = d.AmIHereCol(codiv)
	if err != nil {
		return nil, err
	} else if !t {
		return nil, ErrNotExistCOL
	}
	filter := bson.A{
		bson.D{
			primitive.E{Key: "Date", Value: primitive.Regex{Pattern: `^[0-9]{2}\/[0-9]{2}\/(18|19)$`, Options: ""}},
		},
		bson.D{
			primitive.E{Key: "Index", Value: bson.D{
				primitive.E{Key: "$gte", Value: 1},
			},
			}}}
	var result []string
	collections, err := d.Client.Database(DBDbase).Collection(codiv).Distinct(ctx, "LocalTeam", filter)
	if err != nil {
		return nil, err
	}
	log.Println(collections)
	return result, nil
}

//GetAllCollectionNames ... Get all collection names of an existing database
func (d *DriverMongo) GetAllCollectionNames(db string) ([]string, error) {
	if strings.Trim(db, " ") == "" || len(strings.Trim(db, " ")) == 0 {
		return nil, ErrEmptyString
	}
	str := make([]string, 0)
	ctx := context.Background()
	t, err := d.AmIHereDB(db)
	if err != nil {
		return nil, err
	} else if !t {
		return nil, ErrNotExistDB
	}
	res, err := d.Client.Database(db).ListCollections(ctx, bsonx.Doc{})
	if err != nil {
		return nil, err
	}
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

//AmIHereDB ... Return true if the db exists and false otherwise
func (d *DriverMongo) AmIHereDB(db string) (bool, error) {
	if strings.Trim(db, " ") == "" || len(strings.Trim(db, " ")) == 0 {
		return false, ErrEmptyString
	}
	ctx := context.Background()
	dbs, err := d.Client.ListDatabaseNames(ctx, bson.D{primitive.E{Key: "name", Value: db}})
	if err != nil {
		return false, err
	}
	return len(dbs) == 1, nil
}

//AmIHereCol ... Return true if the col exists and false otherwise
func (d *DriverMongo) AmIHereCol(col string) (bool, error) {
	if strings.Trim(col, " ") == "" || len(strings.Trim(col, " ")) == 0 {
		return false, ErrEmptyString
	}
	str := []string{}
	ctx := context.Background()
	res, err := d.Client.Database(DBDbase).ListCollections(ctx, bson.D{primitive.E{Key: "name", Value: col}})
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

//GetProperties ... Return properties of a team
func (d *DriverMongo) GetProperties(local, away, cols string) (Result, error) {
	if t, err := d.AmIHereDB(DBDbase); err != nil {
		return Result{}, err
	} else if !t {
		return Result{}, nil
	}
	return Result{}, nil
}

//GetStringToConnect ... Return string to connect to sql
func GetStringToConnect() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=admin", DBUser, DBPassword, DBHost, DBPort)
}

//CountryDiv ... Return a map of country:[]divisions
func (d *DriverMongo) CountryDiv() (map[string][]string, error) {
	cols, err := d.GetAllCollectionNames(DBDbase)
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
