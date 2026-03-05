package service

import (
	"context"
	"fmt"
	"go_tweets/internal/config"
	"go_tweets/internal/dto"
	"go_tweets/internal/models"
	"go_tweets/internal/repository"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (string, int, error)
}

type UserServiceImpl struct {
	cfg            *config.Config
	userRepository repository.UserRepository
}

func NewUserService(cfg *config.Config, userRepository repository.UserRepository) UserService {
	return &UserServiceImpl{
		cfg:            cfg,
		userRepository: userRepository,
	}
}

func (s *UserServiceImpl) Register(ctx context.Context, req *dto.RegisterRequest) (string, int, error) {
	userExist, err := s.userRepository.GetUserByEmailOrUsername(ctx, req.Email, req.Username)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	if userExist != nil {
		return "", http.StatusConflict, fmt.Errorf("user %s already exists", req.Username)
	}

	password, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return "", http.StatusUnauthorized, fmt.Errorf("failed to hash password: %v", err)
	}

	userModel := &models.UserModel{
		ID:        uuid.New(),
		Email:     req.Email,
		Username:  req.Username,
		Password:  string(password),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	userId, err := s.userRepository.CreateUser(ctx, userModel)
	if err != nil {
		return "", http.StatusInternalServerError, err
	}

	return userId, http.StatusCreated, nil
}
