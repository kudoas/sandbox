package model

type Todo struct {
	ID     uint   `json:"id" gorm:"primary_key"`
	Text   string `json:"text"`
	Status string `json:"status"`
}
