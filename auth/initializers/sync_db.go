package initializers

import "github.com/TAhirr01/grpc-library/user/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
}
