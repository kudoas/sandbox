package persistance

import (
	"github.com/Kudoas/sandbox/go/simple-api/domain/model"
	"github.com/jinzhu/gorm"
)

type todoRepository struct {
	db *gorm.DB
}

func NewTodoRepository(db *gorm.DB) *todoRepository {
	return &todoRepository{db: db}
}

func (tR *todoRepository) FindAll() ([]*model.Todo, error) {
	todos := []*model.Todo{}
	if err := tR.db.Find(&todos).Error; err != nil {
		return nil, err
	}
	return todos, nil
}
