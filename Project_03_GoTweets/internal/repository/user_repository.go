package repository

import (
	"context"
	"errors"
	"go_tweets/internal/models"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	GetUserByEmailOrUsername(ctx context.Context, email, username string) (*models.UserModel, error)

	CreateUser(ctx context.Context, user *models.UserModel) (string, error)

	GetRefreshToken(ctx context.Context, userId uuid.UUID, now time.Time) (*models.RefreshTokenModel, error)

	StoreRefreshToken(ctx context.Context, refreshToken *models.RefreshTokenModel) error

	GetUserByRefreshToken(ctx context.Context, refreshToken string, now time.Time) (*models.UserModel, *models.RefreshTokenModel, error)

	UpdateRefreshToken(ctx context.Context, id uuid.UUID, newToken string, expiredAt time.Time) error
}

type UserRepositoryImpl struct {
	db *pgxpool.Pool
}

func NewUserRepository(db *pgxpool.Pool) *UserRepositoryImpl {
	return &UserRepositoryImpl{
		db: db,
	}
}

func (repository *UserRepositoryImpl) GetUserByEmailOrUsername(ctx context.Context, email, username string) (*models.UserModel, error) {
	query := `select id, username, email, password, created_at, updated_at from users where email = $1 or username = $2`

	var result models.UserModel

	err := repository.db.QueryRow(ctx, query, email, username).Scan(
		&result.ID,
		&result.Username,
		&result.Email,
		&result.Password,
		&result.CreatedAt,
		&result.UpdatedAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}

func (repository *UserRepositoryImpl) CreateUser(ctx context.Context, user *models.UserModel) (string, error) {
	query := `INSERT INTO users(id, email, username, password, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`

	_, err := repository.db.Exec(ctx, query, user.ID, user.Email, user.Username, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return "", err
	}

	return user.ID.String(), nil
}

func (repository *UserRepositoryImpl) GetRefreshToken(ctx context.Context, userId uuid.UUID, now time.Time) (*models.RefreshTokenModel, error) {
	query := `SELECT id, user_id, refresh_token, expired_at from refresh_tokens WHERE user_id = $1 AND expired_at > $2`

	var result models.RefreshTokenModel
	err := repository.db.QueryRow(ctx, query, userId.String(), now).Scan(
		&result.ID,
		&result.UserId,
		&result.RefreshToken,
		&result.ExpiredAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	return &result, nil
}

func (repository *UserRepositoryImpl) StoreRefreshToken(ctx context.Context, refreshToken *models.RefreshTokenModel) error {
	query := `INSERT INTO refresh_tokens(id, user_id, refresh_token, expired_at, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6)`

	id := uuid.New()
	_, err := repository.db.Exec(ctx, query, id, refreshToken.UserId, refreshToken.RefreshToken, refreshToken.ExpiredAt, refreshToken.CreatedAt, refreshToken.UpdatedAt)
	return err
}

func (repository *UserRepositoryImpl) GetUserByRefreshToken(ctx context.Context, refreshToken string, now time.Time) (*models.UserModel, *models.RefreshTokenModel, error) {
	query := `
		SELECT u.id, u.username, u.email, u.password, u.created_at, u.updated_at,
		       rt.id, rt.user_id, rt.refresh_token, rt.expired_at
		FROM refresh_tokens rt
		INNER JOIN users u ON u.id = rt.user_id
		WHERE rt.refresh_token = $1 AND rt.expired_at > $2
	`

	var user models.UserModel
	var token models.RefreshTokenModel

	err := repository.db.QueryRow(ctx, query, refreshToken, now).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.Password,
		&user.CreatedAt,
		&user.UpdatedAt,
		&token.ID,
		&token.UserId,
		&token.RefreshToken,
		&token.ExpiredAt,
	)

	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, nil, nil
		}
		return nil, nil, err
	}

	return &user, &token, nil
}

func (repository *UserRepositoryImpl) UpdateRefreshToken(ctx context.Context, id uuid.UUID, newToken string, expiredAt time.Time) error {
	query := `UPDATE refresh_tokens SET refresh_token = $1, expired_at = $2, updated_at = $3 WHERE id = $4`

	_, err := repository.db.Exec(ctx, query, newToken, expiredAt, time.Now(), id)
	return err
}
