# Go parameters
GOCMD = go
GOBUILD = $(GOCMD) build
GOCLEAN = $(GOCMD) clean
GOTEST = $(GOCMD) test
GOGET = $(GOCMD) get

# Init environment
init:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest

generate_sql:
	sqlc generate

# Dev target
dev:
	go run . server

# Build target
build:
	$(GOBUILD) -o ./bin/bunny ./cmd

# Clean target
clean:
	$(GOCLEAN)
	rm -rf ./generated
	rm -rf ./bin
