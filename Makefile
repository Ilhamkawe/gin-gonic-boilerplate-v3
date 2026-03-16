.PHONY: run build test docker-up docker-down migrate-up migrate-down

run:
	go run cmd/api/main.go

build:
	go build -o bin/api cmd/api/main.go

test:
	go test -v ./...

docker-up:
	docker-compose up -d

docker-down:
	docker-compose down

# Helper for migrations (example using golang-migrate)
migrate-up:
	migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/warehouse?sslmode=disable" up

migrate-down:
	migrate -path migrations -database "postgres://postgres:postgres@localhost:5432/warehouse?sslmode=disable" down
