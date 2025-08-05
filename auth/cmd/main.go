package main

import (
	"github.com/TAhirr01/grpc-library/auth/initializers"
	"github.com/TAhirr01/grpc-library/auth/pb"
	"github.com/TAhirr01/grpc-library/auth/repo"
	"github.com/TAhirr01/grpc-library/auth/server"
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
	lis, err := net.Listen("tcp", ":50052")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userRepository := repo.NewAuthRepository(db)
	userService := server.NewAuthService(userRepository)
	// Register your service
	pb.RegisterAuthServiceServer(grpcServer, userService)
	log.Println("gRPC Auth server running on :50052")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
