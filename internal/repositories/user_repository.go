package repositories

import (
	"github.com/TAhirr01/grpc-library/internal/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id uint) (*models.User, error)
	GetByIDWithBooks(id uint) (*models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
	AddBookToUser(userID, bookID uint) error
	RemoveBookFromUser(userID, bookID uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) Create(user *models.User) error {
	return r.db.Create(user).Error
}

func (r *userRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := r.db.First(&user, id).Error
	return &user, err
}

func (r *userRepository) GetByIDWithBooks(id uint) (*models.User, error) {
	var user models.User
	err := r.db.Preload("Books").First(&user, id).Error
	return &user, err
}

func (r *userRepository) Update(user *models.User) error {
	return r.db.Save(user).Error
}

func (r *userRepository) Delete(id uint) error {
	return r.db.Delete(&models.User{}, id).Error
}

func (r *userRepository) AddBookToUser(userID, bookID uint) error {
	var user models.User
	var book models.Book
	
	if err := r.db.First(&user, userID).Error; err != nil {
		return err
	}
	
	if err := r.db.First(&book, bookID).Error; err != nil {
		return err
	}
	
	return r.db.Model(&user).Association("Books").Append(&book)
}

func (r *userRepository) RemoveBookFromUser(userID, bookID uint) error {
	var user models.User
	var book models.Book
	
	if err := r.db.First(&user, userID).Error; err != nil {
		return err
	}
	
	if err := r.db.First(&book, bookID).Error; err != nil {
		return err
	}
	
	return r.db.Model(&user).Association("Books").Delete(&book)
}