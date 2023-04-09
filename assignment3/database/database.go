package pkg

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "1234"
	dbname   = "goshop"
)

func GetDB() *gorm.DB {
	dbURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s", user, password, host, port, dbname)
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	if err != nil {
		panic("no connection")
	}
	// db.AutoMigrate(&models.Movie{})
	return db
}