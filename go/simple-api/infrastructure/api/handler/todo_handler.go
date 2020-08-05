package handler

import (
	"net/http"

	"github.com/Kudoas/sandbox/go/simple-api/domain/model"
	"github.com/Kudoas/sandbox/go/simple-api/usecase/repository"
	"github.com/gin-gonic/gin"
)

type todoHandler struct {
	todoRepository repository.TodoRepository
}

type TodoHandler interface {
	GetTodos(c *gin.Context)
}

func NewTodoHandler(tR repository.TodoRepository) TodoHandler {
	return &todoHandler{todoRepository: tR}
}

func (tH *todoHandler) GetTodos(c *gin.Context) {
	e, err := tH.todoRepository.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, e)
}
