package main

import (
	"fmt"
	"database/sql"
	_ "github.com/lib/pq"
)

type User struct {
	Name     string
	Password string
}

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "goshop"
	password = "1234"
)

var (
	db  *sql.DB
	err error
)

func main() {

	db_connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err = sql.Open("postgres", db_connection)

	if err != nil {
		panic(err)
	}

	defer db.Close()

	// insert, err := db.Query("INSERT INTO users (name, password) VALUES ('qqq', '1234')" )
	// if err != nil {
	// 	panic(err)
	// }
	// defer insert.Close()

	res, err := db.Query("SELECT * FROM users ")
	if err != nil {
		panic(err)
	}
	defer res.Close()

	for res.Next() {
		var user User
		err = res.Scan(&user.Name, &user.Password)
		if err != nil {
			panic(err)
		}
		fmt.Printf(fmt.Sprintf("User: %s with password %s", user.Name, user.Password))
	}
}
