package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func Connect() (*pgxpool.Pool, error) {

	connStr := os.Getenv("DB_URL")
	if connStr == "" {
    log.Fatal("DB_URL not set")
}

	pool, err := pgxpool.New(context.Background(), connStr)
	if err != nil {
		return nil, err
	}

	fmt.Println("Connected to postgres")

	return pool, nil

}
