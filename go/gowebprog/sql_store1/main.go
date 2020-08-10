package main

import (
	"log"

	"github.com/jmoiron/sqlx"
	"github.com/tenntenn/sqlite"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

type Diary struct {
	db *sqlx.DB
}

func NewDiary(db *sqlx.DB) *Diary {
	return &Diary{db: db}
}

func (dy *Diary) CreateTable() error {
	sqlStr := `CREATE TABLE IF NOT EXISTS items(
		id        INTEGER PRIMARY KEY,
		content   TEXT NOT NULL,
		author    INTEGER NOT NULL
	)`
	_, err := dy.db.Exec(sqlStr)
	if err != nil {
		return err
	}
	return nil
}

func (dy *Diary) GetPosts() ([]*Post, error) {
	a := make([]*Post, 0)
	sqlStr := `select * from post order by id desc`
	if err := dy.db.Select(&a, sqlStr); err != nil {
		return nil, err
	}
	return a, nil
}

func (dy *Diary) GetPost(id int) (*Post, error) {
	var a *Post
	sqlStr := `select * from post where id = ?`
	if err := dy.db.Get(&a, sqlStr, id); err != nil {
		return nil, err
	}
	return a, nil
}

func (dy *Diary) DeletePost(id int) error {
	if _, err := dy.db.Exec(`delete from post where id = ?`, id); err != nil {
		return err
	}
	return nil
}

func main() {
	db, err := sqlx.Open(sqlite.DriverName, "test.sqlite3")
	defer db.Close()
	if err != nil {
		log.Fatal(err)
	}

	dy := NewDiary(db)
	dy.CreateTable()
}
