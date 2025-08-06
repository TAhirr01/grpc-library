package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name" gorm:"not null"`
	Email string `json:"email" gorm:"unique;not null"`
	Books []Book `gorm:"many2many:user_books;joinForeignKey:UserID;JoinReferences:BookID"`
}