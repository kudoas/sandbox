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
	if err := db.Select(&p, `SELECT * FROM post`); err != nil {
		return nil, err
	}
	return p, nil
}

func GetPost(db *sqlx.DB, id int64) (*model.Post, error) {
	a := model.Post{}
	if err := db.Get(&a, `SELECT * FROM post where id = ?`, id); err != nil {
		return nil, err
	}
	return &a, nil
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

func UpdatePost(db *sqlx.Tx, p *model.Post) (result sql.Result, err error) {
	stmt, err := db.Prepare(`
		UPDATE post SET title = ?, content = ? WHERE id = ?
	`)
	if err != nil {
		return nil, err
	}
	defer func() {
		if closeErr := stmt.Close(); closeErr != nil {
			err = closeErr
		}
	}()
	return stmt.Exec(p.Title, p.Content, p.ID)
}
