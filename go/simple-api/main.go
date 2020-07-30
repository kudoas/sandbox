// refered to this article: https://qiita.com/hyo_07/items/59c093dda143325b1859

package main

import (
	"github.com/Kudoas/sandbox/go/simple-api/controller"
	"github.com/Kudoas/sandbox/go/simple-api/db"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db.ConnectDB()

	router := gin.Default()
	router.GET("/todos", controller.FindTodos)
	router.GET("/todo/:id", controller.FindTodo)
	// router.POST("/todo/new", controller.CreateTodo)
	// router.PATCH("/todo/:id", controller.UpdateTodo)
	// router.DELETE("/todo/:id", controller.DeleteTodo)
	router.Run()
}