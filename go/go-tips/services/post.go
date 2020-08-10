package services

import (
	"example.com/user/go-tips/dbutil"
	"example.com/user/go-tips/model"
	"example.com/user/go-tips/repository"
	"github.com/jmoiron/sqlx"
)

type Post struct {
	db *sqlx.DB
}

func NewPost(db *sqlx.DB) *Post {
	return &Post{db: db}
}

func (p *Post) Create(post *model.Post) (int64, error) {
	var createdId int64
	if err := dbutil.Transact(p.db, func(tx *sqlx.Tx) error {
		result, err := repository.CreatePost(p.db, post)
		if err != nil {
			return err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		createdId = id
		return err
	}); err != nil {
		return 0, err
	}
	return createdId, nil
}
