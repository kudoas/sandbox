package model

import (
	"time"
)

type Todo struct {
	ID       int       `json:"id" gorm:"primary_key"`
	Text     string    `json:"text"`
	Status   string    `json:"status"`
	CreateAt time.Time `json:"createAt"`
}

type InputTodo struct {
	Text   string `json:"text"`
	Status string `json:"status"`
}
