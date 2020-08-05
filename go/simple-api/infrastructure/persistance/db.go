package persistance

import (
	"github.com/Kudoas/sandbox/go/simple-api/domain/model"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
)

type DB struct {
	datasource string
	db         *gorm.DB
}

func NewDB(datasource string) *DB {
	return &DB{datasource: datasource}
}

func (db *DB) Open() (*gorm.DB, error) {
	return gorm.Open("sqlite3", db.datasource)
}

func (db *DB) Migrate(DB *gorm.DB) {
	DB.AutoMigrate(&model.Todo{})
}
