package handlers

import (
	"encoding/json"
	// "golang/assignment3/database"
	"golang/assignment3/models"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
	// "github.com/gin-gonic/gin"
)

type Connection struct {
	DB *gorm.DB
}

func (c *Connection) AddBook(w http.ResponseWriter, r *http.Request) {
	var book []models.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	c.DB.Create(&book)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func (c *Connection) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Query().Get("title")
	books := make([]models.Book, 0)
	if title != "" {
		c.DB.Where("title LIKE ?", title+"%").Find(&books)
	} else {
		c.DB.Find(&books)
	}
	err := json.NewEncoder(w).Encode(books)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (c *Connection) GetBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var book models.Book
	if err := c.DB.First(&book, params["id"]).Error; err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(book)
}