package repository

import (
	"database/sql"

	"example.com/user/go-tips/model"
	"github.com/jmoiron/sqlx"
)

func CreateTable(db *sqlx.DB) error {
	const sqlStr = `CREATE TABLE IF NOT EXISTS post(
		id        INTEGER PRIMARY KEY AUTOINCREMENT,
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

func CreatePost(db *sqlx.DB, p *model.Post) (result sql.Result, err error) {
	stmt, err := db.Prepare(`
		INSERT INTO post (title, content) VALUES (?, ?)
	`)
	if err != nil {
		return nil, err
	}
	result, err = stmt.Exec(p.Title, p.Content)
	if err != nil {
		return nil, err
	}
	return result, nil
}
