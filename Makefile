run:
	@echo "Starting server..."
	@go run cmd/main.go

build:
	@go build -o bin/main cmd/main.go

lint:
	@golangci-lint run ./...