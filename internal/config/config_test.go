package database_test

import (
	"context"
	database "edu_v2/internal/config"
	"github.com/joho/godotenv"
	"testing"
)

func TestConnection(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		return
	}
	ping := database.RDB.Ping(context.TODO())
	if ping != nil {
		t.Fatalf("Failed to connect to Redis: %v", err)
	}
}
