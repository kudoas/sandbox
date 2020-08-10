package controller

import (
	"net/http"

	"example.com/user/go-tips/repository"
	"github.com/jmoiron/sqlx"
)

type Post struct {
	db *sqlx.DB
}

func NewPost(db *sqlx.DB) *Post {
	return &Post{db: db}
}

func (a *Post) Index(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	post, err := repository.AllPost(a.db)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	return http.StatusOK, post, nil
}
