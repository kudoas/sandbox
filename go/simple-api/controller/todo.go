package controller

import (
	"log"
	"net/http"

	"github.com/Kudoas/sandbox/go/simple-api/model"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Todo struct {
	Text   string `json: "title"`
	Status string `json: "status"`
}

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

func Create(c *gin.Context) {
	var todo Todo
	db, err := gorm.Open("sqlite3", "sample-app.sqlite3")
	if err != nil {
		log.Print(err)
	}
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	newTodo := model.Todo{Text: todo.Text, Status: todo.Status}
	db.Create(&newTodo)
	c.JSON(http.StatusOK, gin.H{"data": newTodo})
}

// func Detail(ctx *gin.Context) {
// 	n := ctx.Param("id")
// 	id, err := strconv.Atoi(n)
// 	if err != nil {
// 		log.Print(err)
// 	}
// 	todo := repository.FindTodo(id)
// 	ctx.HTML(200, "detail.html", gin.H{"todo": todo})
// }

// func Update(ctx *gin.Context) {
// 	n := ctx.Param("id")
// 	id, err := strconv.Atoi(n)
// 	if err != nil {
// 		log.Print(err)
// 	}
// 	text := ctx.PostForm("text")
// 	status := ctx.PostForm("status")
// 	repository.UpdateTodo(id, text, status)
// 	ctx.Redirect(302, "/")
// }

// func Delete(ctx *gin.Context) {
// 	n := ctx.Param("id")
// 	id, err := strconv.Atoi(n)
// 	if err != nil {
// 		log.Print(err)
// 	}
// 	repository.DeleteTodo(id)
// 	ctx.Redirect(302, "/")
// }
