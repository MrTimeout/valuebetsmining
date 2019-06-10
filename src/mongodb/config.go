package mongo

import (
	"context"
	"errors"
	"fmt"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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

//GetStringToConnect ... Return string to connect to sql
func GetStringToConnect() string {
	return fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=admin", DBUser, DBPassword, DBHost, DBPort)
}
