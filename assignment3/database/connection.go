package database

import (
	"fmt"
	"golang/assignment3/models"
	"log"
	"time"
	_ "github.com/lib/pq"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func Init() *gorm.DB{
	db, err := gorm.Open("postgres", "user=postgres password=1234 dbname=goshop sslmode=disable")

	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&models.Book{}, &models.User{})

	return db
}

func GetDB() *gorm.DB{
	if db == nil {
		db = Init()
		var sleep = time.Duration(1)
		for db == nil {
			sleep = sleep * 2
			fmt.Print("db is not available, plz wait for 2 seconds", sleep)
			time.Sleep(sleep * time.Second)
			db = Init()
		}
	}

	return db
}