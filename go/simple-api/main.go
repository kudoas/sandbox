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

	r := gin.Default()
	todosGroup := r.Group("/todos")
	todosGroup.GET("", controller.GetTodos)
	todosGroup.GET(":id", controller.GetTodo)
	todosGroup.POST("", controller.CreateTodo)
	todosGroup.PATCH(":id", controller.UpdateTodo)
	// router.DELETE(":id", controller.DeleteTodo)

	r.Run(":8080")
}
