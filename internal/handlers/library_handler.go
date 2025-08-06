package handlers

import (
	"context"
	"github.com/TAhirr01/grpc-library/internal/services"
	"github.com/TAhirr01/grpc-library/pb"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
)

type LibraryHandler struct {
	pb.UnimplementedLibraryServiceServer
	userService services.UserService
	bookService services.BookService
}

func NewLibraryHandler(userService services.UserService, bookService services.BookService) *LibraryHandler {
	return &LibraryHandler{
		userService: userService,
		bookService: bookService,
	}
}

func (h *LibraryHandler) RegisterUser(ctx context.Context, req *pb.RegisterUserRequest) (*pb.UserResponse, error) {
	return h.userService.RegisterUser(req)
}

func (h *LibraryHandler) AddBook(ctx context.Context, req *pb.AddBookRequest) (*pb.BookResponse, error) {
	return h.bookService.AddBook(req)
}

func (h *LibraryHandler) BorrowBook(ctx context.Context, req *pb.BorrowBookRequest) (*pb.BorrowResponse, error) {
	return h.userService.BorrowBook(req)
}

func (h *LibraryHandler) ReturnBook(ctx context.Context, req *pb.BorrowBookRequest) (*pb.BorrowResponse, error) {
	return h.userService.ReturnBook(req)
}

func (h *LibraryHandler) ListUserBooks(ctx context.Context, req *pb.UserID) (*pb.BookListResponse, error) {
	return h.userService.ListUserBooks(uint(req.UserId))
}

func (h *LibraryHandler) ListAllBooks(_ *emptypb.Empty, stream pb.LibraryService_ListAllBooksServer) error {
	books, err := h.bookService.ListAllBooks()
	if err != nil {
		return err
	}

	for _, book := range books {
		if err := stream.Send(&pb.BookResponse{
			Id:     int32(book.ID),
			Title:  book.Title,
			Author: book.Author,
		}); err != nil {
			return err
		}
	}

	return nil
}

func (h *LibraryHandler) UploadBooks(stream pb.LibraryService_UploadBooksServer) error {
	var bookRequests []*pb.AddBookRequest
	
	for {
		bookReq, err := stream.Recv()
		if err == io.EOF {
			// Process all collected books
			summary, err := h.bookService.UploadBooks(bookRequests)
			if err != nil {
				return err
			}
			return stream.SendAndClose(summary)
		}
		if err != nil {
			return err
		}
		
		bookRequests = append(bookRequests, bookReq)
	}
}