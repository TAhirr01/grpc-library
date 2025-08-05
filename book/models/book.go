package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string `json:"book_title"`
	Author string `json:"author"`
}
