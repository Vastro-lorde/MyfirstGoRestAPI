package app

import(
	"golang.org/seun/Todo/app/controllers"
	"github.com/gin-gonic/gin"
)

func StartServer(){
	router := gin.Default()
	router.GET("/todos", controllers.GetTodos)
	router.GET("/", controllers.GetHome)

	router.Run("localhost:5900")
}