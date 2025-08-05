package initializers

import "github.com/TAhirr01/grpc-library/auth/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
