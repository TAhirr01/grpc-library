package interfaces

import "github.com/TAhirr01/grpc-library/user/models"

type UserRepository interface {
	FindAllUsers() ([]models.User, error)
	FindUserById(id uint) (*models.User, error)
	DeleteUserById(id uint) error
	UpdateUser(user *models.User) error
	FindUserByEmail(email string) (*models.User, error)
	CreateUser(user *models.User) error
}
