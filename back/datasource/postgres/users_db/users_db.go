package users_db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"log"
)

var (
	Client   *sql.DB
	user     = "postgres"
	dbname   = "postgres"
	password = "postgres"
)

func init() {
	connStr := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", user, password, dbname)

	Client, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if err := Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully configured")

}
