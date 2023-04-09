package main

import (
	"fmt"
	"golang/assignment3/database"
	"golang/assignment3/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	router := gin.Default()
	
	router.GET("/hello", handlers.Hello)
	router.POST("/registration", handlers.Registration)
	// router.GET("/books", handlers.GetBooks)
	router.POST("/books", handlers.AddBook)
	fmt.Println("Server at 8080")
	router.Run(":8080")
}