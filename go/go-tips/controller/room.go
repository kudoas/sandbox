package controller

import (
	"encoding/json"
	"net/http"

	"example.com/user/go-tips/model"
	"example.com/user/go-tips/repository"
	"example.com/user/go-tips/services"
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

func (a *Post) Create(w http.ResponseWriter, r *http.Request) (int, interface{}, error) {
	newPost := &model.Post{}
	if err := json.NewDecoder(r.Body).Decode(&newPost); err != nil {
		return http.StatusBadRequest, nil, err
	}
	postService := services.NewPost(a.db)
	id, err := postService.Create(newPost)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}
	newPost.ID = id
	return http.StatusOK, newPost, nil
}
