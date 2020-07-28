package repository

import (
	"log"

	"net/http"

	"github.com/Kudoas/sandbox/go/simple-api/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func AllTodo(c *gin.Context) {
	todos := make([]model.Todo, 0)
	db, err := gorm.Open("sqlite3", "sample-app.sqlite3")
	if err != nil {
		log.Print(err)
	}
	db.Find(&todos)
	db.Close()
	c.JSON(http.StatusOK, gin.H{"data": todos})
}

func CreateTodo(text string, status string) {
	db, err := gorm.Open("sqlite3", "sample-app.sqlite3")
	if err != nil {
		log.Println(err)
	}
	db.Create(&model.Todo{Text: text, Status: status})
	db.Close()
}

func UpdateTodo(id int, text string, status string) {
	var todo model.Todo
	db, err := gorm.Open("sqlite3", "sample-app.sqlite3")
	if err != nil {
		log.Print(err)
	}
	db.First(&todo, id)
	todo.Text, todo.Status = text, status
	db.Save(&todo)
	db.Close()
}

func DeleteTodo(id int) {
	var todo model.Todo
	db, err := gorm.Open("sqlite3", "sample-app.sqlite3")
	if err != nil {
		log.Print(err)
	}
	db.First(&todo, id)
	db.Delete(&todo)
	db.Close()
}

func FindTodo(id int) model.Todo {
	var todo model.Todo
	db, err := gorm.Open("sqlite3", "sample-app.sqlite3")
	if err != nil {
		log.Print(err)
	}
	db.First(&todo, id)
	db.Close()
	return todo
}
