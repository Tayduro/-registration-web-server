db:
	docker-compose up -d

migrate-up:
	migrate -source file://./pkg/repository/migrations -database postgres://postgres:12345@localhost:6080/users?sslmode=disable up

migrate-down:
	migrate -source file://./pkg/repository/migrations -database postgres://postgres:12345@localhost:6080/users?sslmode=disable down

build:
	docker-compose up --build

down:
	docker-compose down

start:
	go run ./cmd/signup-server/main.go