// refered to this article: https://qiita.com/hyo_07/items/59c093dda143325b1859

package main

import (
	"log"

	"github.com/Kudoas/sandbox/go/simple-api/infrastructure/api/handler"
	"github.com/Kudoas/sandbox/go/simple-api/infrastructure/persistance"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	newDB := persistance.NewDB("sample-api.sqlite3")
	DB, err := newDB.Open()
	defer DB.Close()
	if err != nil {
		log.Fatal(err)
	}
	if err := newDB.InitDB(); err != nil {
		log.Fatal(err)
	}

	todoRepo := persistance.NewTodoRepository(DB)
	tH := handler.NewTodoHandler(todoRepo)

	r := gin.Default()

	todosGroup := r.Group("/todos")
	todosGroup.GET("", tH.GetTodos)
	// todosGroup.GET(":id", controller.GetTodo)
	// todosGroup.POST("", controller.CreateTodo)
	// todosGroup.PATCH(":id", controller.UpdateTodo)
	// router.DELETE(":id", controller.DeleteTodo)

	r.Run(":8080")
}
