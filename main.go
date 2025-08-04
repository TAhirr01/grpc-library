package main

import (
	"github.com/TAhirr01/grpc-library/initializers"
	"github.com/TAhirr01/grpc-library/pb"
	"github.com/TAhirr01/grpc-library/server"
	"google.golang.org/grpc"
	"log"
	"net"
)

func init() {
	initializers.ConnectToDb()
	initializers.SyncDatabase()
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()

	// Register your service
	pb.RegisterLibraryServiceServer(grpcServer, &server.LibraryServer{})

	log.Println("gRPC Library server running on :50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
