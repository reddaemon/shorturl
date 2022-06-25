package repository

import (
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

type Repo struct {
	db *bolt.DB
}

func NewRepo(db *bolt.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) Set(short string, fullurl string) error {
	err := r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte("Urls"))
		err := b.Put([]byte(short), []byte(fullurl))
		return err
	})

	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) Get(BucketName string, short string) (fullurl []byte, err error) {
	err = r.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketName))
		fullurl = b.Get([]byte(short))
		return nil
	})
	if err != nil {
		return fullurl, err
	}
	if len(fullurl) == 0 {
		return fullurl, fmt.Errorf("fullurl not found: length %v", len(fullurl))
	} 
	return fullurl, nil
}

func (r *Repo) CreateBucket(BucketName string) error {
	err := r.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(BucketName))
		if b == nil {
			log.Printf("creating new bucket: %s", BucketName)
			_, err := tx.CreateBucket([]byte(BucketName))
			if err != nil {
				return fmt.Errorf("create bucket: %s", err)
			}
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}
