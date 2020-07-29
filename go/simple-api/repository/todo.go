package repository

import (
	"log"

	"github.com/Kudoas/sandbox/go/simple-api/model"
	"github.com/jinzhu/gorm"
)

func AllTodo() []model.Todo {
	todos := make([]model.Todo, 0)
	db, err := gorm.Open("sqlite3", "sample-app.sqlite3")
	defer db.Close()
	if err != nil {
		log.Print(err)
	}
	db.Find(&todos)
	return todos
}

func CreateTodo(todo *model.Todo) {
	db, err := gorm.Open("sqlite3", "sample-app.sqlite3")
	defer db.Close()
	if err != nil {
		log.Println(err)
	}
	db.Create(todo)
}

func FindTodo(id int) model.Todo {
	var todo model.Todo
	db, err := gorm.Open("sqlite3", "sample-app.sqlite3")
	defer db.Close()
	if err != nil {
		log.Print(err)
	}
	db.First(&todo, id)
	return todo
}

func UpdateTodo(updateTodo model.Todo) model.Todo {
	var todo model.Todo
	db, err := gorm.Open("sqlite3", "sample-app.sqlite3")
	defer db.Close()
	if err != nil {
		log.Print(err)
	}
	db.First(&todo, updateTodo.ID)
	todo.Text, todo.Status = updateTodo.Text, updateTodo.Status
	db.Save(&todo)
	return todo
}

func DeleteTodo(id int) {
	var todo model.Todo
	db, err := gorm.Open("sqlite3", "sample-app.sqlite3")
	defer db.Close()
	if err != nil {
		log.Print(err)
	}
	db.First(&todo, id)
	db.Delete(&todo)
}
