package repository

import (
	"log"

	"github.com/Kudoas/sandbox/go/simple-api/model"
	"github.com/jinzhu/gorm"
)

func FindAll() ([]*model.Todo, error) {
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

func FindByID(id int) (*model.Todo, error) {
	var todo model.Todo
	db, err := gorm.Open("sqlite3", "sample-app.sqlite3")
	if err != nil {
		return nil, err
	}
	if err := db.First(&todo, id).Error; err != nil {
		return nil, err
	}
	defer db.Close()
	return &todo, nil
}

func Store(todo *model.Todo) error {
	db, err := gorm.Open("sqlite3", "sample-app.sqlite3")
	if err != nil {
		return err
	}
	if err := db.Create(&todo).Error; err != nil {
		return err
	}
	defer db.Close()
	return nil
}

func Update(updateTodo *model.Todo, text string, status string) (*model.Todo, error) {
	var todo model.Todo
	db, err := gorm.Open("sqlite3", "sample-app.sqlite3")
	if err != nil {
		return nil, err
	}
	todo.Text, todo.Status = text, status
	log.Println(todo)
	if err := db.Update(&todo).Error; err != nil {
		return nil, err
	}
	defer db.Close()
	return &todo, nil
}

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
