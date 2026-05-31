# Makefile for Job Application Tracker

.PHONY: run build test clean dev install docker-up docker-down

# Run the application
run:
	go run main.go

# Build the application
build:
	go build -o bin/main main.go

# Run tests
test:
	go test -v ./...

# Clean build artifacts
clean:
	rm -rf bin/
	rm -f main

# Development mode with hot reload (requires air)
dev:
	air

# Install development tools
install-tools:
	go install github.com/cosmtrek/air@latest

# Install dependencies
install:
	go mod download
	go mod tidy

# Run with Docker Compose
docker-up:
	docker-compose up -d

# Stop Docker Compose
docker-down:
	docker-compose down

# View logs
logs:
	docker-compose logs -f

# Database migration (future)
migrate-up:
	@echo "Migration will be added later"

# Format code
fmt:
	go fmt ./...

# Lint code
lint:
	golangci-lint run

# Help
help:
	@echo "Available commands:"
	@echo "  make run          - Run the application"
	@echo "  make build        - Build the application"
	@echo "  make test         - Run tests"
	@echo "  make clean        - Clean build artifacts"
	@echo "  make dev          - Run with hot reload"
	@echo "  make install      - Install dependencies"
	@echo "  make docker-up    - Start with Docker Compose"
	@echo "  make docker-down  - Stop Docker Compose"
	@echo "  make fmt          - Format code"
	@echo "  make help         - Show this help"
