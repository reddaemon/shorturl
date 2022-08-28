package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

func InitDB() *sql.DB {
	db, err := sql.Open("sqlite3", "urls")
	if err != nil {
		log.Fatal(err)
	}

	return db
}
