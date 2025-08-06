package services

import (
	"errors"
	"github.com/TAhirr01/grpc-library/internal/models"
	"github.com/TAhirr01/grpc-library/internal/repositories"
	"github.com/TAhirr01/grpc-library/pb"
)

type UserService interface {
	RegisterUser(req *pb.RegisterUserRequest) (*pb.UserResponse, error)
	BorrowBook(req *pb.BorrowBookRequest) (*pb.BorrowResponse, error)
	ReturnBook(req *pb.BorrowBookRequest) (*pb.BorrowResponse, error)
	ListUserBooks(userID uint) (*pb.BookListResponse, error)
}

type userService struct {
	userRepo repositories.UserRepository
	bookRepo repositories.BookRepository
}

func NewUserService(userRepo repositories.UserRepository, bookRepo repositories.BookRepository) UserService {
	return &userService{
		userRepo: userRepo,
		bookRepo: bookRepo,
	}
}

func (s *userService) RegisterUser(req *pb.RegisterUserRequest) (*pb.UserResponse, error) {
	user := &models.User{
		Name:  req.Name,
		Email: req.Email,
	}
	
	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}
	
	return &pb.UserResponse{
		Id:    int32(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}

func (s *userService) BorrowBook(req *pb.BorrowBookRequest) (*pb.BorrowResponse, error) {
	// Check if user exists
	user, err := s.userRepo.GetByIDWithBooks(uint(req.UserId))
	if err != nil {
		return nil, errors.New("user not found")
	}
	
	// Check if book exists
	book, err := s.bookRepo.GetByID(uint(req.BookId))
	if err != nil {
		return nil, errors.New("book not found")
	}
	
	// Check if user already has this book
	for _, userBook := range user.Books {
		if userBook.ID == book.ID {
			return nil, errors.New("book already assigned to user")
		}
	}
	
	// Add book to user
	if err := s.userRepo.AddBookToUser(uint(req.UserId), uint(req.BookId)); err != nil {
		return nil, errors.New("failed to borrow book")
	}
	
	return &pb.BorrowResponse{Message: "Book borrowed successfully"}, nil
}

func (s *userService) ReturnBook(req *pb.BorrowBookRequest) (*pb.BorrowResponse, error) {
	// Check if user exists
	if _, err := s.userRepo.GetByID(uint(req.UserId)); err != nil {
		return nil, errors.New("user not found")
	}
	
	// Check if book exists
	if _, err := s.bookRepo.GetByID(uint(req.BookId)); err != nil {
		return nil, errors.New("book not found")
	}
	
	// Remove book from user
	if err := s.userRepo.RemoveBookFromUser(uint(req.UserId), uint(req.BookId)); err != nil {
		return nil, errors.New("failed to return book")
	}
	
	return &pb.BorrowResponse{Message: "Book returned successfully"}, nil
}

func (s *userService) ListUserBooks(userID uint) (*pb.BookListResponse, error) {
	user, err := s.userRepo.GetByIDWithBooks(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}
	
	var bookResponses []*pb.BookResponse
	for _, book := range user.Books {
		bookResponses = append(bookResponses, &pb.BookResponse{
			Id:     int32(book.ID),
			Title:  book.Title,
			Author: book.Author,
		})
	}
	
	return &pb.BookListResponse{Books: bookResponses}, nil
}