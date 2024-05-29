FROM golang:1.22-alpine AS build

WORKDIR /app
COPY . .

RUN go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN sqlc generate
RUN swag init -g main.go

RUN go mod download
RUN go mod verify
RUN go mod tidy

RUN go build -o megacrm .

FROM golang:1.22-alpine

WORKDIR /app
COPY --from=build /app/megacrm .
COPY .env.example .env
CMD [ "./bunny", "server" ]
