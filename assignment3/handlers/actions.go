package handlers

import (
	"encoding/json"
	"golang/assignment3/database"
	"golang/assignment3/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Hello(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"response": gin.H{
			"method":  http.MethodGet,
			"code":    http.StatusOK,
			"message": "Hello world",
		},
	})
}

func Registration(context *gin.Context) {
	var user *models.User

	decode := json.NewDecoder(context.Request.Body).Decode(&user)

	if decode != nil {
		context.JSON(http.StatusOK, gin.H{
			"response": decode.Error(),
		})
		return
	}

	isExist := database.IsExistUserByName(user.Name)

	if !isExist {
		user := &models.User{Name: user.Name, Email: user.Email, Password: user.Password}
		isSuccessAdded := database.Add(user)
		if isSuccessAdded == nil {
			context.JSON(http.StatusOK, gin.H{
				"response": gin.H{
					"code":    http.StatusOK,
					"message": "You are successfully registered",
				},
			})
		}
	} else {
		context.JSON(http.StatusOK, gin.H{
			"response": "user with this name already exists",
		})
		return
	}

}

var db *gorm.DB

func GetBooks(c *gin.Context) {
	var books []models.Book
	db.Find(&books)
	c.JSON(http.StatusOK, books)
}

func AddBook(c *gin.Context) {
	// parse the JSON request body into an Body struct
	var book models.Book
	err := c.BindJSON(&book)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// save the book to the database
	result := db.Create(&book)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	// return the newly created book
	c.JSON(http.StatusOK, book)
}



// func AddBook(c *gin.Context) {
// 	var book []models.Book
// 	if err := c.BindJSON(&book); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	db.Create(&book)
// 	c.JSON(http.StatusCreated, book)
// }
