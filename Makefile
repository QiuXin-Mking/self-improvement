# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=train
BINARY_WEB=web_app
DOCKER_IMAGE_NAME=spaced-repetition-go

# Build the project
build-cli:
	$(GOBUILD) -o bin/$(BINARY_NAME) -v main.go

build-web:
	$(GOBUILD) -o bin/$(BINARY_WEB) -v web_server.go

build: build-cli build-web

# Run the project
run-cli:
	go run main.go

run-web:
	JWT_SECRET=your-secret-key-for-development go run web_server.go

# Install dependencies
deps:
	go mod tidy
	cd frontend && npm install

# Clean the project
clean:
	$(GOCLEAN)
	rm -f bin/$(BINARY_NAME)
	rm -f bin/$(BINARY_WEB)

# Initialize knowledge base
init:
	go run main.go --init

# Show statistics
stats:
	go run main.go --stats

# Test the application
test:
	$(GOTEST) -v ./...

# Build frontend
frontend-build:
	cd frontend && npm run build

# Run frontend dev server
frontend-dev:
	cd frontend && npm run dev

# Database setup
db-init:
	@echo "Initializing database with migrations..."
	@mkdir -p data
	@sqlite3 data/app.db < migrations/001_initial_schema.sql
	@echo "Database initialized at data/app.db"

# Run with database initialization
run-web-with-db: db-init run-web

.PHONY: build build-cli build-web run-cli run-web deps clean init stats test frontend-build frontend-dev db-init run-web-with-db