package repository

import (
	"context"
	"database/sql"
	"errors"
	"go_tweets/internal/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

type UserRepository interface {
	GetUserByEmailOrUsername(ctx context.Context, email, username string) (*models.UserModel, error)

	CreateUser(ctx context.Context, user *models.UserModel) (string, error)
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &result, nil
}

func (repository *UserRepositoryImpl) CreateUser(ctx context.Context, user *models.UserModel) (string, error) {
	query := `INSERT INTO user(id, email, username, password) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id, email, username, created_at, updated_at`

	id := uuid.New()
	_, err := repository.db.Exec(ctx, query, id, user.Email, user.Username, user.Password, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return "", err
	}

	return id.String(), nil

}
