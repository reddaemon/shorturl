package repository

import (
	"database/sql"
	"github.com/pkg/errors"
	"time"
)

type Repository interface {
	Set(short string, fullUrl string) (id int64, err error)
	Get(short string) (fullUrl string, err error)
}

type store struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return &store{db: db}
}

func (r *store) Set(short string, fullUrl string) (id int64, err error) {
	query := "INSERT INTO urls(shortUrl,fullUrl,created) VALUES(?,?,?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return id, errors.Wrapf(err, "cannot prepare query %s", query)
	}
	res, err := stmt.Exec(short, fullUrl, time.Now())
	if err != nil {
		return id, errors.Wrapf(err, "cannot insert, short %s, full %s", short, fullUrl)
	}
	id, err = res.LastInsertId()
	if err != nil {
		return id, errors.Wrapf(err, "cannot get last inserted id")
	}
	return id, nil
}

func (r *store) Get(short string) (fullUrl string, err error) {

	query := "SELECT fullUrl FROM urls WHERE shortUrl = ?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return fullUrl, errors.Wrapf(err, "cannot prepare query %s", query)
	}
	res := stmt.QueryRow(short)
	if err != nil {
		return fullUrl, errors.Wrapf(err, "cannot get result %s", short)
	}

	err = res.Scan(&fullUrl)
	if err != nil {
		return fullUrl, errors.Wrapf(err, "cannot scan result")
	}

	if fullUrl == "" {
		return fullUrl, errors.New("full url not found")
	}

	return fullUrl, nil
}
