package db

import (
	"log"

	"github.com/Kudoas/sandbox/go/simple-api/model"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

func ConnectDB() {
	dbName := "sample-app.sqlite3"
	db, err := gorm.Open("sqlite3", dbName)
	defer db.Close()
	if err != nil {
		log.Print(err)
	}
	db.AutoMigrate(&model.Todo{})
}
