package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"golang.org/seun/Todo/app/data"
)

func GetTodos(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, data.TodoList)
}
