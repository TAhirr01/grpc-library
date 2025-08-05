package interfaces

import "github.com/TAhirr01/grpc-library/auth/models"

type AuthRepository interface {
	CreateUser(user *models.User) error
	FindUserById(id uint) (*models.User, error)
	FindUserByEmail(email string) (*models.User, error)
}
