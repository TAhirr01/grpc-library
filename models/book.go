package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title  string `json:"book_title" gorm:"not null"`
	Author string `json:"author" gorm:"not null"`
	Users  []User `gorm:"many2many:user_books;joinForeignKey:UserID;JoinReferences:BookID"`
}
