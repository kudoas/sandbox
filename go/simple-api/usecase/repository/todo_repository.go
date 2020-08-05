package repository

import (
	"github.com/Kudoas/sandbox/go/simple-api/domain/model"
)

type TodoRepository interface {
	FindAll() ([]*model.Todo, error)
}
