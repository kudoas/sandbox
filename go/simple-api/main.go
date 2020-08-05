// refered to this article: https://qiita.com/hyo_07/items/59c093dda143325b1859

package main

import (
	"fmt"
	"log"

	"github.com/Kudoas/sandbox/go/simple-api/infrastructure/api/handler"
	"github.com/Kudoas/sandbox/go/simple-api/infrastructure/persistance"
	"github.com/jinzhu/gorm"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func InitDB(datasource string) (*gorm.DB, error) {
	newDB := persistance.NewDB("sample-api.splite3")
	DB, err := newDB.Open()
	if err != nil {
		return nil, fmt.Errorf("failed db init. %s", err)
	}
	newDB.Migrate(DB)
	return DB, nil
}

func main() {
	DB, err := InitDB("sample-app.sqlite3")
	defer DB.Close()
	if err != nil {
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
