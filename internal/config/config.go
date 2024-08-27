package database

import (
	"context"
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"log"
	"os"
)

var DB *sql.DB
var RDB *redis.Client

func ConnectPostgres() {
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=%s host=%s port=%s",
		os.Getenv("DO_POSTGRES_USER"),
		os.Getenv("DO_POSTGRES_PASSWORD"),
		os.Getenv("DO_POSTGRES_DATABASE"),
		os.Getenv("DO_POSTGRES_SSLMODE"),
		os.Getenv("DO_POSTGRES_HOST"),
		os.Getenv("DO_POSTGRES_PORT"),
	)

	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database connection: ", err)
	}
	if err := DB.Ping(); err != nil {
		log.Fatalf("Error pinging database: %v", err)
	}

	fmt.Println("Database connection successful")
}

func ConnectRedis() {
	redisHost := os.Getenv("REDIS_HOST")
	redisPort := os.Getenv("REDIS_PORT")

	RDB = redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", redisHost, redisPort),
	})

	ctx := context.Background()
	_, err := RDB.Ping(ctx).Result()
	if err != nil {
		fmt.Printf("Failed to connect to Redis: %v\n", err)
		return
	}
	fmt.Println("Redis connection established successfully")
}
