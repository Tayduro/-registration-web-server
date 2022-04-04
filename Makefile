#build:
#	docker-compose build
#	docker-compose up postgres
#
#run:
#	docker run -v /mnt/c/Work/registration-web-server/pkg/repository/migrations:/migrations --network host migrate/migrate -path=/migrations/ -database "postgres://postgres:12345@localhost:6080/users?sslmode=disable" up
#	docker-compose up
#
#stop:
#	docker-compose down
#
db:
	docker-compose up -d

migrate-up:
	migrate -source file://./pkg/repository/migrations -database postgres://postgres:12345@localhost:6080/users?sslmode=disable up

migrate-down:
	migrate -source file://./pkg/repository/migrations -database postgres://postgres:12345@localhost:6080/users?sslmode=disable down

run:
	docker-compose up --build

down:
	docker-compose down