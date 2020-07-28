// refered to this article: https://qiita.com/hyo_07/items/59c093dda143325b1859

package main

import (
	"github.com/Kudoas/sandbox/go/simple-api/db"
	"github.com/Kudoas/sandbox/go/simple-api/repository"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db.ConnectDB()

	router := gin.Default()
	router.GET("/todos", repository.AllTodo)
	// router.POST("/new", controller.Create)
	// router.GET("/:id", controller.Detail)
	// router.PATCH("/:id", controller.Update)
	// router.DELETE("/:id", controller.Delete)
	router.Run()
}
