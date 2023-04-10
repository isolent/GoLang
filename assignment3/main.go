package main

import (
	"fmt"
	"log"
	"os"
	"strconv"

	// "golang/assignment3/database"
	"golang/assignment3/handlers"
	"golang/assignment3/models"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)



func main() {

	err := godotenv.Load()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Almaty",
		os.Getenv("DB_HOST"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_NAME"), os.Getenv("DB_PORT"))

	// Parse port to integer
	_, err = strconv.Atoi(os.Getenv("DB_PORT"))
	if err != nil {
		log.Fatalf("Error parsing DB_PORT: %v", err)
	}

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
	router.HandleFunc("/search/", c.SearchBookByTitle).Methods("GET")
	router.HandleFunc("/sorted-books/", c.GetSortedBooks).Methods("GET")

	fmt.Println("Server at 8181")
	http.ListenAndServe(":8181", router)
}