package main

import (
	"database/sql"
	"log"
)

func connect(url string) *sql.DB {
	database, err := sql.Open("postgres", url)
	if err != nil {
		log.Fatalln("Could not connect to db", err.Error())
	}
	return database
}
