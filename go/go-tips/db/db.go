package db

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	datasource string
}

func NewDB(datasource string) *DB {
	return &DB{datasource: datasource}
}

func (db *DB) Open() (*sqlx.DB, error) {
	return sqlx.Open("sqlite3", db.datasource)
}
