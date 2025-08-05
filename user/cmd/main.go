package main

import (
	"github.com/TAhirr01/grpc-library/user/initializers"
	"github.com/TAhirr01/grpc-library/user/pb"
	"github.com/TAhirr01/grpc-library/user/repo"
	"github.com/TAhirr01/grpc-library/user/server"
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
	lis, err := net.Listen("tcp", ":50055")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userRepository := repo.NewUserRepository(db)
	userService := server.NewUserService(userRepository)
	// Register your service
	pb.RegisterUserServiceServer(grpcServer, userService)
	log.Println("gRPC User server running on :50055")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
