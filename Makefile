# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get

# Init environment
init:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install github.com/air-verse/air@latest
	go install github.com/swaggo/swag/cmd/swag@latest
	go install go.uber.org/nilaway/cmd/nilaway@latest
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest

	cp ./pre-commit .git/hooks/pre-commit
	chmod +x .git/hooks/pre-commit
generate_sql:
	sqlc generate

generate_openapi:
	swag init -g main.go

# Dev target
server-dev:
	go run . server

# Run example consumer: make consumer name=example
.PHONY: consumer
consumer:
	go run . consumer $(name)

# Build target
build:
	$(GOBUILD) -o ./bin/bunny .

migrate-up:
	migrate -database postgresql://bunny:bunny@localhost:5432/bunny?sslmode=disable -path db/migrations up

migrate-down:
	migrate -database postgresql://bunny:bunny@localhost:5432/bunny?sslmode=disable -path db/migrations down 1

migrate-create:
	migrate create -ext sql -dir db/migrations -tz utc $(name)

# Clean target
clean:
	$(GOCLEAN)
	rm -rf ./generated
	rm -rf ./bin

lint:
	golangci-lint run .

lint-fix:
	golangci-lint run --fix .

lint-changes:
	@changed_files=$$(git status --porcelain | grep -E '^(M|A|R|C)' | awk '{print $$2}' | grep '\.go$$'); \
	if [ -n "$$changed_files" ]; then \
		echo "Running golangci-lint on changed files..."; \
		golangci-lint run $$changed_files; \
	else \
		echo "No Go files changed. Skipping lint."; \
	fi

lint-fix-changes:
	@changed_files=$$(git status --porcelain | grep -E '^(M|A|R|C)' | awk '{print $$2}' | grep '\.go$$'); \
	if [ -n "$$changed_files" ]; then \
		echo "Running golangci-lint --fix on changed files..."; \
		golangci-lint run --fix $$changed_files; \
	else \
		echo "No Go files changed. Skipping lint."; \
	fi