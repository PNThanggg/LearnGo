package repository

import (
	"context"
	"fmt"
	"time"
	"todo-apps/internal/models"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgxpool"
)

func CreateToDo(pool *pgxpool.Pool, title string, completed bool) (*models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	id := uuid.New()
	var query = "INSERT INTO todos (id, title, completed) VALUES ($1, $2, $3) RETURNING id, title, completed, created_at, updated_at"

	var todo models.Todo

	err := pool.QueryRow(ctx, query, id, title, completed).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func GetAllTodos(pool *pgxpool.Pool) ([]models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	rows, err := pool.Query(ctx, "SELECT * FROM todos ORDER BY created_at DESC")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var todos []models.Todo
	for rows.Next() {
		var todo models.Todo

		err = rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Completed,
			&todo.CreatedAt,
			&todo.UpdatedAt,
		)

		if err != nil {
			return nil, err
		}

		todos = append(todos, todo)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}

	return todos, nil
}

func GetTodoById(pool *pgxpool.Pool, id string) (*models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var query = "SELECT * FROM todos WHERE id = $1"

	var todo models.Todo

	err := pool.QueryRow(ctx, query, id).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func UpdateTodo(pool *pgxpool.Pool, id string, title string, completed bool) (*models.Todo, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var query = "UPDATE todos SET title = $1, completed = $2, update_at = CURRENT_TIMESTAMP WHERE id = $3 returning id, title, completed, created_at,  updated_at"

	var todo models.Todo
	err := pool.QueryRow(ctx, query, title, completed, id).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Completed,
		&todo.CreatedAt,
		&todo.UpdatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &todo, nil
}

func DeleteTodo(pool *pgxpool.Pool, id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var query = "DELETE FROM todos WHERE id = $1"
	result, err := pool.Exec(ctx, query, id)
	if err != nil {
		return err
	}

	if result.RowsAffected() == 0 {
		return fmt.Errorf("todo with id %s does not exist", id)
	}

	return nil
}
