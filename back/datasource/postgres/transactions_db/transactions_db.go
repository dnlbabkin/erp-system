package transactions_db

import (
	"database/sql"
	"fmt"
	"log"
)

var (
	Client   *sql.DB
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	dbname   = "transactions"
	password = "postgres"
)

func InitDB() *sql.DB {
	connStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	Client, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}

	if err := Client.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully configured")

	return Client
}
