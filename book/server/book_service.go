package server

import (
	"context"
	"errors"
	"github.com/TAhirr01/grpc-library/book/models"
	"github.com/TAhirr01/grpc-library/book/pb"
	"github.com/TAhirr01/grpc-library/book/repo/interfaces"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
)

type BookService struct {
	pb.UnimplementedBookServiceServer
	repo interfaces.BookRepository
}

func NewBookService(repo interfaces.BookRepository) *BookService {
	return &BookService{repo: repo}
}

func (b *BookService) AddBook(ctx context.Context, req *pb.AddBookRequest) (*pb.BookResponse, error) {
	book := models.Book{Title: req.Title, Author: req.Author}
	if err := b.repo.CreateBook(&book); err != nil {
		return nil, err
	}
	return &pb.BookResponse{Id: int32(book.ID), Title: book.Title, Author: book.Author}, nil
}

func (b *BookService) DeleteBook(ctx context.Context, req *pb.DeleteBookRequest) (*emptypb.Empty, error) {
	book, err := b.repo.FindBookById(uint(req.Id))
	if err != nil {
		return nil, err
	}
	if book == nil {
		return nil, errors.New("book not found")
	}
	if err := b.repo.DeleteBookById(uint(req.Id)); err != nil {
		return nil, errors.New("book delete failed")
	}
	return nil, nil
}

func (b *BookService) ListAllBooks(_ *emptypb.Empty, stream pb.BookService_ListAllBooksServer) error {
	books, err := b.repo.FindAllBooks()
	if err != nil {
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

func (b *BookService) UploadBooks(stream pb.BookService_UploadBooksServer) error {
	count := 0
	for { //sonsuz loop kitablarin gelmeyini temin edir
		bookReq, err := stream.Recv() //1 AddBookRequest gelir bolck olur sora tezeden burdan basdiyir
		if err == io.EOF {            //Client streami dayandirib gondermek qurtarib
			return stream.SendAndClose(&pb.UploadSummary{ //Geriye upload olan kitablarin sayni qaytarir sorada sterami dayanidir rpc qutarir
				Count: int32(count),
			})
		}
		if err != nil { //Basqa case errorlar ola biler meselen client disconnect olub
			return err
		}

		book := models.Book{
			Title:  bookReq.Title,
			Author: bookReq.Author,
		}
		if err := b.repo.CreateBook(&book); err != nil {
			return err
		}
		count++
	}
}
