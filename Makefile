# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get

# Init environment
init:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
	go install github.com/cosmtrek/air@latest

generate_sql:
	sqlc generate

generate_openapi:
	swag init -g main.go

# Dev target
server-dev:
	go run . server

example-consumer-dev:
	go run . consumer example

# Build target
build:
	$(GOBUILD) -o ./bin/bunny .

# Clean target
clean:
	$(GOCLEAN)
	rm -rf ./generated
	rm -rf ./bin
