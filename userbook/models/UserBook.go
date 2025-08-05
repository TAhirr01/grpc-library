package models

import "gorm.io/gorm"

type UserBook struct {
	gorm.Model
	BookID uint
	UserID uint
}
