package controller

import (
	"net/http"
	"strconv"
	"time"

	"github.com/Kudoas/sandbox/go/simple-api/model"
	"github.com/Kudoas/sandbox/go/simple-api/repository"
	"github.com/gin-gonic/gin"
)

func GetTodos(c *gin.Context) {
	todos, err := repository.FindAll()
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": todos})
}

func GetTodo(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		c.JSON(http.StatusBadRequest, model.ResponseError{Message: err.Error()})
	}
	todo, err := repository.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, model.ResponseError{Message: err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"todo": todo})
}

func CreateTodo(c *gin.Context) {
	var todo model.InputTodo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if todo.Text == "" || todo.Status == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Null input data"})
		return
	}
	newTodo := model.Todo{Text: todo.Text, Status: todo.Status, CreateAt: time.Now()}
	if err := repository.Store(&newTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": newTodo})
}

func UpdateTodo(c *gin.Context) {
	n := c.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	todo, err := repository.FindByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err})
		return
	}
	var inputTodo model.InputTodo
	if err := c.ShouldBindJSON(&inputTodo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if _, err := repository.Update(todo, inputTodo.Text, inputTodo.Status); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": todo})
}

// func DeleteTodo(c *gin.Context) {
// 	n := c.Param("id")
// 	id, err := strconv.Atoi(n)
// 	if err != nil {
// 		log.Print(err)
// 	}
// 	repository.DeleteTodo(id)
// 	c.JSON(http.StatusOK, nil)
// }
