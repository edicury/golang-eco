package main

import (
	_ "github.com/lib/pq"
	"database/sql"
	"log"
)


func getDB() *sql.DB {
	connStr := "postgres://postgres:123quatro@localhost/goeco?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	return db
}

