package main

import (
	"github.com/TAhirr01/grpc-library/book/initializers"
	"github.com/TAhirr01/grpc-library/book/pb"
	repo "github.com/TAhirr01/grpc-library/book/repo"
	"github.com/TAhirr01/grpc-library/book/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func init() {
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	db := initializers.DB
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	bookRepository := repo.NewBookRepository(db)
	bookService := server.NewBookService(bookRepository)
	// Register your service
	pb.RegisterBookServiceServer(grpcServer, bookService)
	log.Println("gRPC Book server running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
