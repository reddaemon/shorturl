package repository

import (
	"database/sql/driver"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

var setTests = []struct {
	short string
	full  string
}{
	{"http://localhost:8080/Jr", "google.com"},
	{"http://localhost:8080/Gk", "youtube.com"},
	{"http://localhost:8080/Sd", "mail.ru"},
}

var getTests = []struct {
	id    int64
	short string
	want  string
}{
	{1, "http://localhost:8080/Jr", "https://youtube.com"},
	{2, "http://localhost:8080/Gk", "https://google.com"},
	{3, "http://localhost:8080/Sd", "https://mail.ru"},
}

func TestSet(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func() {
		_ = db.Close()
	}()

	repo := NewRepository(db)

	for i, e := range setTests {
		query := "INSERT INTO urls(shortUrl,fullUrl,created) VALUES(?,?,?)"
		mock.ExpectPrepare(query).ExpectExec().WithArgs(e.short, e.full, AnyTime{}).WillReturnResult(sqlmock.NewResult(int64(i), 1))
		id, err := repo.Set(e.short, e.full)
		if err != nil {
			t.Fatal(err)
		}
		assert.NoError(t, err)
		assert.Equal(t, int64(i), id, "test id")

	}

}

func TestGet(t *testing.T) {
	db, mock, err := sqlmock.New(sqlmock.QueryMatcherOption(sqlmock.QueryMatcherEqual))
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer func() {
		_ = db.Close()
	}()

	repo := NewRepository(db)

	for _, e := range getTests {
		query := "SELECT fullUrl FROM urls WHERE shortUrl = ?"
		rows := sqlmock.NewRows([]string{"fullUrl"}).
			AddRow(e.want)
		mock.ExpectPrepare(query).ExpectQuery().WithArgs(e.short).WillReturnRows(rows)
		fullUrl, err := repo.Get(e.short)
		if err != nil {
			t.Fatal(err)
		}
		assert.NoError(t, err)
		assert.Equal(t, e.want, fullUrl, "test fullUrl")
	}

}

type AnyTime struct{}

func (a AnyTime) Match(v driver.Value) bool {
	_, ok := v.(time.Time)
	return ok
}
