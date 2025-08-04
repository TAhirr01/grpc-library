package initializers

import "github.com/TAhirr01/grpc-library/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Book{})
}
