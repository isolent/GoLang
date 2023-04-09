package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	_ "github.com/lib/pq"
)

// func homePage(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Hello")
// }

func main() {

	r := gin.Default()
	r.GET("/books", GetBooks)
	fmt.Println("Server at 8080")
	http.ListenAndServe(":8080", nil)
	
}

var db *gorm.DB

func GetBooks(c *gin.Context) {

	var books [] Book
	db.Find(&books)
	c.JSON(http.StatusOK, books)
}

type Book struct {
	gorm.Model
	Title       string  `gorm:"unique"`
	Description string  `json:"description"`
	Cost        float32 `json:"cost"`
}