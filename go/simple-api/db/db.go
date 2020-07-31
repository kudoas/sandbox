package db

import (
	"github.com/Kudoas/sandbox/go/simple-api/model"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

var DB *gorm.DB

func ConnectDB() {
	database, err := gorm.Open("sqlite3", "sample-app.sqlite3")
	if err != nil {
		panic("Failed to cennect to database")
	}
	database.AutoMigrate(&model.Todo{})
	DB = database
}
