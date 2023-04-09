package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string  `gorm:"unique"`
	Description string  `json:"description"`
	Cost        float32 `json:"cost"`
}

type User struct {
	gorm.Model
	Name        string  `gorm:"unique"`
	Email 		string  `gorm:"unique"`
	Password  	string  `gorm:"unique"`
}