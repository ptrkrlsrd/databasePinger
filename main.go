package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

func main() {
	var connStr string
	flag.StringVar(&connStr, "connection", "", "Connection string")
	flag.Parse()

	if len(os.Args) == 1 {
		flag.Usage()
	}

	if connStr == "" {
		log.Fatal("No connection string. Please provide one using the '-connection <string>' flag")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("PING")
}
