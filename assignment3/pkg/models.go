package pkg
import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title       string  `gorm:"unique"`
	Description string  `json:"description"`
	Cost        float32 `json:"cost"`
}
