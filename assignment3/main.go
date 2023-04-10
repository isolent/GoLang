package main

import (
	"fmt"
	"golang/handlers"
	"golang/models"
	"log"
	"net/http"
	// "strconv"

	"github.com/gorilla/mux"
	// "github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {

	// err := godotenv.Load()

	// dsn := "host=database user=postgres password=1234 dbname=assign3 port=5433 sslmode=disable"
	dsn := "host=localhost user=postgres password=1234 dbname=assign3 port=5432 sslmode=disable"

	db, _ := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err := db.AutoMigrate(&models.Book{}); err != nil {
		log.Fatal(err)
	}

	c := handlers.Connection{DB: db}

	router := mux.NewRouter()

	router.HandleFunc("/books/", c.GetAllBooks).Methods("GET")
	router.HandleFunc("/books/{id}/", c.GetBookByID).Methods("GET")
	router.HandleFunc("/addbook/", c.AddBook).Methods("POST")
	router.HandleFunc("/updatebooks/{id}/", c.UpdateBook).Methods("PUT")
	router.HandleFunc("/deletebooks/{id}/", c.DeleteBookByID).Methods("DELETE")
	router.HandleFunc("/search/{title}/", c.SearchBookByTitle).Methods("GET")
	router.HandleFunc("/sortedBooks/", c.GetSortedBooks).Methods("GET")
	router.HandleFunc("/sortedBooksDesc/", c.DescGetSortedBooks).Methods("GET")

	fmt.Println("Server at 8181")
	http.ListenAndServe(":8181", router)
}
