package repository

import (
	"example.com/user/go-tips/model"
	"github.com/jmoiron/sqlx"
)

func CreateTable(db *sqlx.DB) error {
	const sqlStr = `CREATE TABLE IF NOT EXISTS post(
		id        INTEGER PRIMARY KEY,
		title     TEXT NOT NULL,
		content   INTEGER NOT NULL
	);`
	_, err := db.Exec(sqlStr)
	return err
}

func AllPost(db *sqlx.DB) ([]model.Post, error) {
	p := make([]model.Post, 0)
	if err := db.Select(&p, `SELECT id, title FROM post`); err != nil {
		return nil, err
	}
	return p, nil
}
