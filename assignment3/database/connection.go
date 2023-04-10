package database

import (
	"golang/assignment3/models"
	"log"
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Init() *gorm.DB{
	db, err := gorm.Open("postgres", "user=postgres password=1234 dbname=goshop sslmode=disable")

	if err != nil {
		log.Fatal(err)
		// panic(err)
	}

	db.AutoMigrate(&models.Book{})

	return db
}
