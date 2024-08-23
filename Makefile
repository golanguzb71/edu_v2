include .env
DB_USER := $(POSTGRES_USER)
DB_PASS := $(POSTGRES_PASSWORD)
DB_NAME := $(POSTGRES_DATABASE)
DB_HOST := $(POSTGRES_HOST)
DB_PORT := $(POSTGRES_PORT)

DATABASE_URL := postgres://$(DB_USER):$(DB_PASS)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable

.SILENT:

run:
	go run cmd/main.go

migrate-up:
	@echo "Running migrations up..."
	psql $(DATABASE_URL) -f migrations/up.sql

migrate-down:
	@echo "Running migrations down..."
	psql $(DATABASE_URL) -f migrations/down.sql

migrate-mock:
	@echo "Running mock migrations..."
	psql $(DATABASE_URL) -f migrations/mock_information.sql
