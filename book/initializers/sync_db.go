package initializers

import "github.com/TAhirr01/grpc-library/book/models"

func SyncDatabase() {
	DB.AutoMigrate(&models.Book{})
}
