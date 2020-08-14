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
		if err := tx.Commit(); err != nil {
			return err
		}
		id, err := result.LastInsertId()
		if err != nil {
			return err
		}
		createdId = id
		return nil
	}); err != nil {
		return 0, err
	}
	return createdId, nil
}

func (p *Post) Show(id int64) (*model.Post, error) {
	showPost := &model.Post{}
	if err := dbutil.Transact(p.db, func(tx *sqlx.Tx) error {
		post, err := repository.GetPost(p.db, id)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		showPost = post
		return nil
	}); err != nil {
		return nil, err
	}
	return showPost, nil
}

func (p *Post) Update(post *model.Post) error {
	if err := dbutil.Transact(p.db, func(tx *sqlx.Tx) error {
		if _, err := repository.UpdatePost(tx, post); err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		return nil
	}); err != nil {
		return err
	}
	return nil
}
