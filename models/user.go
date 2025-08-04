package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name  string `json:"name"`
	Email string `json:"email" gorm:"unique"`
	Books []Book `gorm:"many2many:user_books;joinForeignKey:BookID;JoinReferences:UserID"`
}
