package services

import (
	"github.com/TAhirr01/grpc-library/internal/models"
	"github.com/TAhirr01/grpc-library/internal/repositories"
	"github.com/TAhirr01/grpc-library/pb"
)

type BookService interface {
	AddBook(req *pb.AddBookRequest) (*pb.BookResponse, error)
	ListAllBooks() ([]models.Book, error)
	UploadBooks(books []*pb.AddBookRequest) (*pb.UploadSummary, error)
}

type bookService struct {
	bookRepo repositories.BookRepository
}

func NewBookService(bookRepo repositories.BookRepository) BookService {
	return &bookService{
		bookRepo: bookRepo,
	}
}

func (s *bookService) AddBook(req *pb.AddBookRequest) (*pb.BookResponse, error) {
	book := &models.Book{
		Title:  req.Title,
		Author: req.Author,
	}
	
	if err := s.bookRepo.Create(book); err != nil {
		return nil, err
	}
	
	return &pb.BookResponse{
		Id:     int32(book.ID),
		Title:  book.Title,
		Author: book.Author,
	}, nil
}

func (s *bookService) ListAllBooks() ([]models.Book, error) {
	return s.bookRepo.GetAll()
}

func (s *bookService) UploadBooks(bookRequests []*pb.AddBookRequest) (*pb.UploadSummary, error) {
	var books []models.Book
	for _, req := range bookRequests {
		books = append(books, models.Book{
			Title:  req.Title,
			Author: req.Author,
		})
	}
	
	if err := s.bookRepo.CreateBatch(books); err != nil {
		return nil, err
	}
	
	return &pb.UploadSummary{Count: int32(len(books))}, nil
}