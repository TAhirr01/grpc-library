package repositories

import (
	"github.com/TAhirr01/grpc-library/internal/models"
	"gorm.io/gorm"
)

type BookRepository interface {
	Create(book *models.Book) error
	GetByID(id uint) (*models.Book, error)
	GetAll() ([]models.Book, error)
	Update(book *models.Book) error
	Delete(id uint) error
	CreateBatch(books []models.Book) error
}

type bookRepository struct {
	db *gorm.DB
}

func NewBookRepository(db *gorm.DB) BookRepository {
	return &bookRepository{db: db}
}

func (r *bookRepository) Create(book *models.Book) error {
	return r.db.Create(book).Error
}

func (r *bookRepository) GetByID(id uint) (*models.Book, error) {
	var book models.Book
	err := r.db.First(&book, id).Error
	return &book, err
}

func (r *bookRepository) GetAll() ([]models.Book, error) {
	var books []models.Book
	err := r.db.Find(&books).Error
	return books, err
}

func (r *bookRepository) Update(book *models.Book) error {
	return r.db.Save(book).Error
}

func (r *bookRepository) Delete(id uint) error {
	return r.db.Delete(&models.Book{}, id).Error
}

func (r *bookRepository) CreateBatch(books []models.Book) error {
	return r.db.CreateInBatches(books, 100).Error
}