package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	connStr := os.Getenv("POSTGRES_CONNECTION")
	if connStr == "" {
		log.Fatal("No connection string. Please provide one using the $POSTGRES_CONNECTION environment variable")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}
