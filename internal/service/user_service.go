package service

import (
	"context"
	"log"
	"time"

	UserProto "github.com/JacobRWebb/InventoryManagement.Users.Api/pkg/api"
	"github.com/JacobRWebb/InventoryManagement/internal/client"
	"github.com/JacobRWebb/InventoryManagement/internal/models"
)

type userService struct {
	clients *client.ProtoClients
}

func NewUserService(clients *client.ProtoClients) UserService {
	s := &userService{
		clients: clients,
	}

	return s
}

func (s *userService) RegisterUser(email string, password string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	registerUserRequest := &UserProto.RegisterUserRequest{
		Email:    email,
		Password: password,
	}

	_, err := s.clients.UserServiceClient.RegisterUser(ctx, registerUserRequest)

	if err != nil {
		return err
	}

	return nil
}

func (s *userService) LoginUser(email string, password string) (*models.AuthResponse, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	loginUserRequest := &UserProto.LoginUserRequest{
		Email:    email,
		Password: password,
	}

	response, err := s.clients.UserServiceClient.LoginUser(ctx, loginUserRequest)

	if err != nil {
		return nil, err
	}

	authRes := &models.AuthResponse{
		AccessToken:  response.AccessToken,
		RefreshToken: response.RefreshToken,
		TokenType:    response.TokenType,
	}

	return authRes, nil
}

func (s *userService) ValidateToken(token string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	validateTokenRequest := &UserProto.ValidateTokenRequest{
		AccessToken: token,
	}

	response, err := s.clients.UserServiceClient.ValidateToken(ctx, validateTokenRequest)

	if err != nil {
		return "", err
	}

	return response.UserId, nil
}

func (s *userService) GetUser(userId string) (*models.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	getUserRequest := &UserProto.GetUserRequest{
		UserId: userId,
	}

	getUserResponse, err := s.clients.UserServiceClient.GetUser(ctx, getUserRequest)

	if err != nil {
		return nil, err
	}

	user := &models.User{
		Id:    getUserResponse.Id,
		Email: getUserResponse.Email,
		Profile: &models.UserProfile{
			FirstName:          getUserResponse.Profile.FirstName,
			ProfilePic:         getUserResponse.Profile.AvatarUrl,
			PhoneNumber:        "",
			OnboardingComplete: getUserResponse.Profile.FullName == "",
		},
	}
	log.Printf("%v", user)
	return user, nil
}
