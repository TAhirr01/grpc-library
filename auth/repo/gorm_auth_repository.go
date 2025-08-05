package repo

import (
	"github.com/TAhirr01/grpc-library/auth/models"
	"github.com/TAhirr01/grpc-library/auth/repo/interfaces"
	"gorm.io/gorm"
)

type gormAuthRepository struct {
	db *gorm.DB
}

func NewAuthRepository(db *gorm.DB) interfaces.AuthRepository {
	return &gormAuthRepository{db: db}
}

func (ur *gormAuthRepository) FindUserById(id uint) (*models.User, error) {
	var user models.User
	err := ur.db.First(&user, id).Error
	return &user, err
}

func (ur *gormAuthRepository) CreateUser(user *models.User) error {
	return ur.db.Create(user).Error
}

func (ur *gormAuthRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := ur.db.Where("email = ?", email).First(&user).Error
	return &user, err
}
