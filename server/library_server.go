package server

import (
	"context"
	"errors"
	"github.com/TAhirr01/grpc-library/initializers"
	"github.com/TAhirr01/grpc-library/models"
	"github.com/TAhirr01/grpc-library/pb"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
)

type LibraryServer struct {
	pb.UnimplementedLibraryServiceServer
}

func (s *LibraryServer) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.UserResponse, error) {
	user := models.User{Name: req.Name, Email: req.Email}
	if err := initializers.DB.Create(&user).Error; err != nil {
		return nil, err
	}
	return &pb.UserResponse{Id: int32(user.ID), Name: user.Name, Email: user.Email}, nil
}

func (s *LibraryServer) AddBook(ctx context.Context, req *pb.AddBookRequest) (*pb.BookResponse, error) {
	book := models.Book{Title: req.Title, Author: req.Author}
	if err := initializers.DB.Create(&book).Error; err != nil {
		return nil, err
	}
	return &pb.BookResponse{Id: int32(book.ID), Title: book.Title, Author: book.Author}, nil
}

func (s *LibraryServer) BorrowBook(ctx context.Context, req *pb.BorrowBookRequest) (*pb.BorrowResponse, error) {
	var user models.User
	var book models.Book
	if err := initializers.DB.Preload("Books").First(&user, req.UserId).Error; err != nil {
		return nil, errors.New("user not found")
	}
	if err := initializers.DB.First(&book, req.BookId).Error; err != nil {
		return nil, errors.New("book not found")
	}

	for _, b := range user.Books {
		if b.ID == book.ID {
			return nil, errors.New("book already assigned to user")
		}
	}

	if err := initializers.DB.Model(&user).Association("Books").Append(&book); err != nil {
		return nil, errors.New("cannot add new book")
	}
	return &pb.BorrowResponse{Message: "Book added successfully"}, nil
}

func (s *LibraryServer) ReturnBook(ctx context.Context, req *pb.BorrowBookRequest) (*pb.BorrowResponse, error) {
	var user models.User
	var book models.Book
	if err := initializers.DB.First(&user, req.UserId).Error; err != nil {
		return &pb.BorrowResponse{Message: "user not found"}, errors.New("user not found")
	}
	if err := initializers.DB.First(&book, req.BookId).Error; err != nil {
		return &pb.BorrowResponse{Message: "book not found"}, errors.New("book not found")
	}
	if err := initializers.DB.Model(&user).Association("Books").Delete(&book); err != nil {
		return &pb.BorrowResponse{Message: "cannot delete book"}, errors.New(" Cannot delete book")
	}
	return &pb.BorrowResponse{Message: "Book removed successfully"}, nil
}

func (s *LibraryServer) ListUserBooks(ctx context.Context, req *pb.UserID) (*pb.BookListResponse, error) {
	var user models.User
	if err := initializers.DB.Preload("Books").First(&user, req.UserId).Error; err != nil {
		return &pb.BookListResponse{Books: nil}, errors.New("user not found")
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
func (s *LibraryServer) ListAllBooks(_ *emptypb.Empty, stream pb.LibraryService_ListAllBooksServer) error {
	var books []models.Book
	if err := initializers.DB.Find(&books).Error; err != nil {
		return err
	}

	for _, book := range books {
		if err := stream.Send(&pb.BookResponse{
			Id:     int32(book.ID),
			Title:  book.Title,
			Author: book.Author,
		}); err != nil {
			return err // exit if client closes connection
		}
	}

	return nil
}
func (s *LibraryServer) UploadBooks(stream pb.LibraryService_UploadBooksServer) error {
	count := 0
	for {
		bookReq, err := stream.Recv()
		if err == io.EOF {
			// All books received, send summary
			return stream.SendAndClose(&pb.UploadSummary{
				Count: int32(count),
			})
		}
		if err != nil {
			return err
		}

		book := models.Book{
			Title:  bookReq.Title,
			Author: bookReq.Author,
		}
		if err := initializers.DB.Create(&book).Error; err != nil {
			return err
		}
		count++
	}
}
