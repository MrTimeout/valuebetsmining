package postgres

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

	// Check that our connection is good
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return &Db{db}, nil
}

//postgresql://[user[:password]@][netloc][:port][,...][/dbname][?param1=value1&...]
func ConnString(host string, port int, user string, dbName string, password string) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%d/%s?sslmode=disabled",
		user, password, host, port, dbName,
	)
}

type User struct {
	ID         int
	Name       string
	Age        int
	Profession string
	Friendly   bool
}

func (d *Db) GetUsersByName(name string) []User {

	stmt, err := d.Prepare("SELECT * FROM users WHERE name=$1")
	if err != nil {
		fmt.Println("GetUserByName Preperation Err: ", err)
	}

	rows, err := stmt.Query(name)
	if err != nil {
		fmt.Println("GetUserByName Query Err: ", err)
	}

	var r User
	users := []User{}

	for rows.Next() {
		err = rows.Scan(
			&r.ID,
			&r.Name,
			&r.Age,
			&r.Profession,
			&r.Friendly,
		)
		if err != nil {
			fmt.Println("Error scanning rows: ", err)
		}
		users = append(users, r)
	}

	return users
}