package interfaces

import "github.com/TAhirr01/grpc-library/book/models"

type BookRepository interface {
	CreateBook(book *models.Book) error
	FindAllBooks() ([]models.Book, error)
	FindBookById(bookID uint) (*models.Book, error)
	DeleteBookById(bookID uint) error
	UpdateBook(book *models.Book) error
}
