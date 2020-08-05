package persistance

import (
	"fmt"

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

// DB接続
func (db *DB) Open() (*gorm.DB, error) {
	return gorm.Open("sqlite3", db.datasource)
}

// migrate
func (db *DB) migrate(DB *gorm.DB) {
	DB.AutoMigrate(&model.Todo{})
}

// init
func (db *DB) InitDB() error {
	DB, err := db.Open()
	if err != nil {
		return fmt.Errorf("failed db init. %s", err)
	}
	db.migrate(DB)
	return nil
}
