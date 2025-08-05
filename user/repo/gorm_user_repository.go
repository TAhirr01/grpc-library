package repo

import (
	"github.com/TAhirr01/grpc-library/user/models"
	"github.com/TAhirr01/grpc-library/user/repo/interfaces"
	"gorm.io/gorm"
)

type gormUserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) interfaces.UserRepository {
	return &gormUserRepository{db: db}
}

func (ur *gormUserRepository) FindAllUsers() ([]models.User, error) {
	var users []models.User
	err := ur.db.Find(&users).Error
	return users, err
}
func (ur *gormUserRepository) FindUserById(id uint) (*models.User, error) {
	var user models.User
	err := ur.db.First(&user, id).Error
	return &user, err
}
func (ur *gormUserRepository) DeleteUserById(id uint) error {
	return ur.db.Delete(&models.User{}, id).Error
}

func (ur *gormUserRepository) UpdateUser(user *models.User) error {
	return ur.db.Save(user).Error
}

func (ur *gormUserRepository) FindUserByEmail(email string) (*models.User, error) {
	var user models.User
	err := ur.db.Where("email = ?", email).First(&user).Error
	return &user, err
}

func (ur *gormUserRepository) CreateUser(user *models.User) error {
	return ur.db.Create(user).Error
}
