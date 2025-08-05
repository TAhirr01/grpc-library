package server

import (
	"github.com/TAhirr01/grpc-library/user/pb"
	"github.com/TAhirr01/grpc-library/user/repo/interfaces"
)

type UserService struct {
	userRepository interfaces.UserRepository
	pb.UnimplementedUserServiceServer
}

func NewUserService(userRepository interfaces.UserRepository) *UserService {
	return &UserService{userRepository: userRepository}
}
