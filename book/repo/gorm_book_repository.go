package repo

import (
	"github.com/TAhirr01/grpc-library/book/models"
	"github.com/TAhirr01/grpc-library/book/repo/interfaces"
	"gorm.io/gorm"
)

type gormBookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) interfaces.BookRepository {
	return &gormBookRepository{db: db}
}

func (repo *gormBookRepository) CreateBook(book *models.Book) error {
	return repo.db.Create(book).Error
}

func (repo *gormBookRepository) FindAllBooks() ([]models.Book, error) {
	var books []models.Book
	err := repo.db.Find(&books).Error
	return books, err
}
func (repo *gormBookRepository) DeleteBookById(bookID uint) error {
	return repo.db.Delete(&models.Book{}, bookID).Error
}

func (repo *gormBookRepository) FindBookById(bookID uint) (*models.Book, error) {
	var book models.Book
	err := repo.db.First(&book, bookID).Error
	return &book, err
}

func (repo *gormBookRepository) UpdateBook(book *models.Book) error {
	return repo.db.Save(book).Error
}
