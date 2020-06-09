package models

import (
	"github.com/jinzhu/gorm"
)

// Book ...
type Book struct {
	gorm.Model
	Name     string `json:"name"`
	Author   string `json:"author"`
	Category string `json:"category"`
	Price    int    `json:"price"`
}

// TableName ...
func (b *Book) TableName() string {
	return "book"
}
