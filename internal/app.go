package internal

import (
	"log"
	"os"
	"github.com/gin-gonic/gin"
	"golang.org/seun/Todo/internal/controllers"
)

func StartServer() {
	router := gin.Default()
	router.GET("/todos", controllers.GetTodos)
	router.GET("/", controllers.GetHome)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	if err := router.Run(":" + port); err != nil {
		log.Panicf("error: %s", err)
	}
}
