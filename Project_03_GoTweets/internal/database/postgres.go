package database

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
)

func ConnectDb(databaseUrl string) (*pgxpool.Pool, error) {
	var ctx = context.Background()

	config, err := pgxpool.ParseConfig(databaseUrl)
	if err != nil {
		log.Printf("Unable to parse database URL: %v", err)
		return nil, err
	}

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		log.Printf("Unable to create connection pool: %v", err)
		return nil, err
	}

	err = pool.Ping(ctx)
	if err != nil {
		log.Printf("Unable to ping database: %v", err)
		pool.Close()
		return nil, err
	}

	log.Printf("Successfully connected to PostgreSQL database")

	return pool, nil
}
