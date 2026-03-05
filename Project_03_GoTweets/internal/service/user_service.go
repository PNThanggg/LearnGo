package service

import (
	"context"
	"fmt"
	"go_tweets/internal/config"
	"go_tweets/internal/dto"
	"go_tweets/internal/models"
	"go_tweets/internal/repository"
	pkgJwt "go_tweets/pkg/jwt"
	"go_tweets/refresh_token"
	"net/http"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(ctx context.Context, req *dto.RegisterRequest) (string, int, error)

	Login(ctx context.Context, req *dto.LoginRequest) (string, string, int, error)
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
		return "", http.StatusInternalServerError, fmt.Errorf("failed to hash password: %v", err)
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

func (s *UserServiceImpl) Login(ctx context.Context, req *dto.LoginRequest) (string, string, int, error) {
	userExist, err := s.userRepository.GetUserByEmailOrUsername(ctx, "", req.Username)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	if userExist == nil {
		return "", "", http.StatusUnauthorized, fmt.Errorf("user %s does not exist", req.Username)
	}

	err = bcrypt.CompareHashAndPassword([]byte(userExist.Password), []byte(req.Password))
	if err != nil {
		return "", "", http.StatusUnauthorized, fmt.Errorf("password incorrect")
	}

	token, err := pkgJwt.CreateToken(userExist.ID, userExist.Username, s.cfg.JwtSecret)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	now := time.Now()
	refreshTokenModel, err := s.userRepository.GetRefreshToken(ctx, userExist.ID, now)
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	if refreshTokenModel != nil {
		return token, refreshTokenModel.RefreshToken, http.StatusOK, nil
	}

	refreshToken, err := refresh_token.GenerateRefreshToken()
	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	err = s.userRepository.StoreRefreshToken(
		ctx, &models.RefreshTokenModel{
			UserId:       userExist.ID,
			RefreshToken: refreshToken,
			ExpiredAt:    time.Now().Add(7 * 24 * time.Hour),
			CreatedAt:    time.Now(),
			UpdatedAt:    time.Now(),
		})

	if err != nil {
		return "", "", http.StatusInternalServerError, err
	}

	return token, refreshToken, http.StatusOK, nil
}
