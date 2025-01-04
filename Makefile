run:
	@echo "Starting server..."
	@go run dev/main.go

build:
	@go build -o bin/main dev/main.go

lint:
	@golangci-lint run ./...

deploy:
	@echo "Deploying to Vercel (preview)..."
	vercel

deploy-prod:
	@echo "Deploying to Vercel..."
	vercel --prod