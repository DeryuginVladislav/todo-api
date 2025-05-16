package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func InitDB() {
	dbUrl := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	var err error
	DB, err = pgxpool.New(context.Background(), dbUrl)
	if err != nil {
		panic(fmt.Sprintf("Unable to connect to database: %v\n", err))
	}

	// Проверка подключения
	err = DB.Ping(context.Background())
	if err != nil {
		panic(fmt.Sprintf("Unable to ping database: %v\n", err))
	}

	fmt.Println("Successfully connected to database!")
}

func CloseDB() {
	DB.Close()
}
