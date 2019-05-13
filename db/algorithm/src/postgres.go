package algorithm

import (
	"database/sql"
	"fmt"

	// postgres driver
	_ "github.com/lib/pq"
)

type Db struct {
	*sql.DB
}

func New(connString string) (*Db, error) {
	db, err := sql.Open("postgres", connString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Db{db}, nil
}

//postgresql://[user[:password]@][netloc][:port][,...][/dbname][?param1=value1&...]
func ConnString(host string, port string, user string, dbName string, password string) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		user, password, host, port, dbName,
	)
}