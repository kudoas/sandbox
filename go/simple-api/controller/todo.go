package controller

import (
	"log"
	"strconv"

	"github.com/Kudoas/sandbox/go/simple-api/repository"
	"github.com/gin-gonic/gin"
)

// func Index(ctx *gin.Context) {
// 	repository.AllTodo()
// }

func Create(ctx *gin.Context) {
	text := ctx.PostForm("text")
	status := ctx.PostForm("status")
	repository.CreateTodo(text, status)
	ctx.Redirect(302, "/")
}

func Detail(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		log.Print(err)
	}
	todo := repository.FindTodo(id)
	ctx.HTML(200, "detail.html", gin.H{"todo": todo})
}

func Update(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		log.Print(err)
	}
	text := ctx.PostForm("text")
	status := ctx.PostForm("status")
	repository.UpdateTodo(id, text, status)
	ctx.Redirect(302, "/")
}

func Delete(ctx *gin.Context) {
	n := ctx.Param("id")
	id, err := strconv.Atoi(n)
	if err != nil {
		log.Print(err)
	}
	repository.DeleteTodo(id)
	ctx.Redirect(302, "/")
}
