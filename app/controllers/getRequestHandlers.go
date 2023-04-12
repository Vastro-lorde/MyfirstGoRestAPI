package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/seun/Todo/app/data"
)

func GetTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, data.TodoList)
}

func GetHome(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, "Welcome to My first go rest api")
}
