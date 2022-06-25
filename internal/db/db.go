package db

import (
	"log"

	"github.com/boltdb/bolt"
)

func InitDB() *bolt.DB {
	db, err := bolt.Open("urls.db", 0600, nil)
	if err != nil {
		log.Fatal(err)
	}

	return db
}
