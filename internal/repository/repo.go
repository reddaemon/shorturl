package repository

import (
	"database/sql"
	"errors"
	"time"
)

type RepoTool interface {
	Set(short string, fullUrl string) (id int64, err error)
	Get(short string) (fullUrl string, err error)
}

type Repo struct {
	db *sql.DB
}

func NewRepo(db *sql.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) Set(short string, fullUrl string) (id int64, err error) {
	query := "INSERT INTO urls(shortUrl,fullUrl,created) VALUES(?,?,?)"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return id, err
	}
	res, err := stmt.Exec(short, fullUrl, time.Now())
	if err != nil {
		return id, err
	}
	id, err = res.LastInsertId()
	if err != nil {
		return id, err
	}
	return id, nil
}

func (r *Repo) Get(short string) (fullUrl string, err error) {

	query := "SELECT fullUrl FROM urls WHERE shortUrl = ?"
	stmt, err := r.db.Prepare(query)
	if err != nil {
		return fullUrl, err
	}
	res := stmt.QueryRow(short)
	if err != nil {
		return fullUrl, err
	}

	err = res.Scan(&fullUrl)
	if err != nil {
		return fullUrl, err
	}

	if fullUrl == "" {
		return fullUrl, errors.New("full url not found")
	}

	return fullUrl, nil
}
