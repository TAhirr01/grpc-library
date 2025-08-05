package server

import (
	"context"
	"errors"
	"github.com/TAhirr01/grpc-library/auth/models"
	"github.com/TAhirr01/grpc-library/auth/pb"
	"github.com/TAhirr01/grpc-library/auth/repo/interfaces"
	"github.com/TAhirr01/grpc-library/auth/utils"
	"golang.org/x/crypto/bcrypt"
	"log"
	"time"
)

type AuthService struct {
	pb.UnimplementedAuthServiceServer
	repo interfaces.AuthRepository
}

func NewAuthService(repo interfaces.AuthRepository) *AuthService {
	return &AuthService{repo: repo}
}

func (a *AuthService) RegisterUser(ctx context.Context, request *pb.RegisterUserRequest) (*pb.UserResponse, error) {
	log.Println("Someone tries to register a new user")
	user, err := a.repo.FindUserByEmail(request.Email)
	if err == nil && user != nil {
		log.Println("User already exists")
		return nil, errors.New("user with this email already exists")
	}
	hash, err := utils.HashPassword(request.Password)
	if err != nil {
		log.Println("Failed to hash the password")
		return nil, errors.New("failed to create user")
	}
	newUser := models.User{
		Name:         request.Name,
		Email:        request.Email,
		PasswordHash: hash,
	}
	log.Println("User created successfully")
	if err := a.repo.CreateUser(&newUser); err != nil {
		return nil, errors.New("failed to create user")
	}

	return &pb.UserResponse{
		Id:    int32(newUser.ID),
		Name:  newUser.Name,
		Email: newUser.Email,
	}, nil
}

func (a *AuthService) LoginUser(ctx context.Context, req *pb.LoginRequest) (*pb.LoginResponse, error) {
	log.Println("Someone tries to log in")
	user, err := a.repo.FindUserByEmail(req.Email)
	// Step 1: Find user by email
	if err != nil {
		return nil, errors.New("user not found")
	}

	// Step 2: Compare password
	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Step 3: Generate JWT token
	token, err := utils.GenerateToken(user.ID, 15*time.Minute)
	if err != nil {
		log.Println("Failed to generate access token")
		return nil, errors.New("failed to generate token")
	}

	// Step 4: Return token
	return &pb.LoginResponse{
		Token: token,
	}, nil
}

func (a *AuthService) ValidateToken(ctx context.Context, req *pb.ValidateRequest) (*pb.UserData, error) {
	userID, err := utils.ValidateJWT(req.Token)
	if err != nil {
		return nil, errors.New("unauthorized: invalid token")
	}
	user, err := a.repo.FindUserById(userID)
	if err != nil {
		return nil, errors.New("user not found")
	}

	return &pb.UserData{
		Id:    int32(user.ID),
		Name:  user.Name,
		Email: user.Email,
	}, nil
}
