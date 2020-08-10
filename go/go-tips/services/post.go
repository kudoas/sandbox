package services

import (
	"github.com/jmoiron/sqlx"
)

type Post struct {
	db *sqlx.DB
}

func NewPost(db *sqlx.DB) *Post {
	return &Post{db: db}
}

// func (p *Post) Create(post *model.Post) (int64, error) {
// 	var createdId int64

// }
