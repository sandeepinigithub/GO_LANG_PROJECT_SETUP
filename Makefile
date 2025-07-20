# Variables
APP_NAME=devsmailgo-api
BUILD_DIR=bin
DOCKER_IMAGE=devsmailgo-api
DOCKER_TAG=latest

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
BINARY_NAME=$(APP_NAME)
BINARY_UNIX=$(BINARY_NAME)_unix

# Default target
.DEFAULT_GOAL := help

.PHONY: help
help: ## Show this help message
	@echo 'Usage: make [target]'
	@echo ''
	@echo 'Targets:'
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "  %-15s %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: build
build: ## Build the application
	$(GOBUILD) -o $(BUILD_DIR)/$(BINARY_NAME) cmd/main.go

.PHONY: run
run: ## Run the application
	$(GOCMD) run cmd/main.go

.PHONY: test
test: ## Run tests
	$(GOTEST) -v ./...

.PHONY: test-coverage
test-coverage: ## Run tests with coverage
	$(GOTEST) -v -cover ./...

.PHONY: clean
clean: ## Clean build artifacts
	$(GOCLEAN)
	rm -rf $(BUILD_DIR)

.PHONY: deps
deps: ## Download dependencies
	$(GOMOD) download

.PHONY: deps-update
deps-update: ## Update dependencies
	$(GOMOD) get -u ./...

.PHONY: deps-tidy
deps-tidy: ## Tidy dependencies
	$(GOMOD) mod tidy

.PHONY: lint
lint: ## Run linter
	golangci-lint run

.PHONY: format
format: ## Format code
	gofmt -s -w .

.PHONY: docker-build
docker-build: ## Build Docker image
	docker build -t $(DOCKER_IMAGE):$(DOCKER_TAG) .

.PHONY: docker-run
docker-run: ## Run Docker container
	docker run -p 8080:8080 --env-file .env $(DOCKER_IMAGE):$(DOCKER_TAG)

.PHONY: docker-compose-up
docker-compose-up: ## Start services with Docker Compose
	docker-compose up -d

.PHONY: docker-compose-down
docker-compose-down: ## Stop services with Docker Compose
	docker-compose down

.PHONY: docker-compose-logs
docker-compose-logs: ## View Docker Compose logs
	docker-compose logs -f

.PHONY: docker-compose-restart
docker-compose-restart: ## Restart services with Docker Compose
	docker-compose restart

.PHONY: migrate
migrate: ## Run database migrations
	$(GOCMD) run cmd/main.go

.PHONY: seed
seed: ## Seed database with sample data
	$(GOCMD) run scripts/seed.go

.PHONY: health
health: ## Check application health
	curl -f http://localhost:8080/api/health || exit 1

.PHONY: install-tools
install-tools: ## Install development tools
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	go install github.com/go-delve/delve/cmd/dlv@latest

.PHONY: dev-setup
dev-setup: ## Setup development environment
	cp env.example .env
	$(GOMOD) download
	@echo "Development environment setup complete!"
	@echo "Please edit .env file with your configuration."

.PHONY: release
release: ## Build release binary
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -a -installsuffix cgo -o $(BUILD_DIR)/$(BINARY_UNIX) cmd/main.go

.PHONY: security-scan
security-scan: ## Run security scan
	gosec ./...

.PHONY: benchmark
benchmark: ## Run benchmarks
	$(GOTEST) -bench=. ./...

.PHONY: generate
generate: ## Generate code (if using code generation)
	$(GOCMD) generate ./...

.PHONY: swagger
swagger: ## Generate Swagger documentation
	swag init -g cmd/main.go

.PHONY: proto
proto: ## Generate protobuf code
	protoc --go_out=. --go_opt=paths=source_relative \
		--go-grpc_out=. --go-grpc_opt=paths=source_relative \
		proto/*.proto

.PHONY: clean-all
clean-all: clean ## Clean everything including Docker
	docker system prune -f
	docker volume prune -f

.PHONY: logs
logs: ## View application logs
	tail -f logs/app.log

.PHONY: monitor
monitor: ## Monitor application resources
	@echo "CPU Usage:"
	top -p $$(pgrep $(BINARY_NAME) | head -1) -n 1
	@echo "Memory Usage:"
	ps aux | grep $(BINARY_NAME) | grep -v grep

.PHONY: backup
backup: ## Backup database
	mysqldump -u devsmailgo -p devsmailgo > backup/devsmailgo_$(shell date +%Y%m%d_%H%M%S).sql

.PHONY: restore
restore: ## Restore database from backup
	@echo "Usage: make restore BACKUP_FILE=backup/filename.sql"
	@if [ -z "$(BACKUP_FILE)" ]; then echo "Please specify BACKUP_FILE"; exit 1; fi
	mysql -u devsmailgo -p devsmailgo < $(BACKUP_FILE) 