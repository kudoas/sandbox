package repository

import (
	"log"

	"github.com/Kudoas/sandbox/go/simple-api/model"
	"github.com/jinzhu/gorm"
)

func AllTodo() ([]*model.Todo, error) {
	todos := []*model.Todo{}
	db, err := gorm.Open("sqlite3", "sample-app.sqlite3")
	if err != nil {
		return nil, err
	}
	if err := db.Find(&todos).Error; err != nil {
		return nil, err
	}
	defer db.Close()
	return todos, nil
}

func FindTodo(id int) (*model.Todo, error) {
	var todo model.Todo
	db, err := gorm.Open("sqlite3", "sample-app.sqlite3")
	if err != nil {
		log.Print(err)
	}
	if err := db.First(&todo, id).Error; err != nil {
		return nil, err
	}
	defer db.Close()
	return &todo, nil
}

// func CreateTodo(todo *model.Todo) error {
// 	db, err := gorm.Open("sqlite3", "sample-app.sqlite3")
// 	if err != nil {
// 		log.Println(err)
// 	}
// 	defer db.Close()
// 	return db.Create(&todo).Error
// }

// func UpdateTodo(updateTodo model.Todo) model.Todo {
// 	var todo model.Todo
// 	db, err := gorm.Open("sqlite3", "sample-app.sqlite3")
// 	if err != nil {
// 		log.Print(err)
// 	}
// 	db.First(&todo, updateTodo.ID)
// 	todo.Text, todo.Status = updateTodo.Text, updateTodo.Status
// 	db.Save(&todo)
// 	defer db.Close()
// 	return todo
// }

// func DeleteTodo(id int) {
// 	var todo model.Todo
// 	db, err := gorm.Open("sqlite3", "sample-app.sqlite3")
// 	if err != nil {
// 		log.Print(err)
// 	}
// 	db.First(&todo, id)
// 	db.Delete(&todo)
// 	defer db.Close()
// }