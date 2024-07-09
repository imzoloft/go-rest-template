build:
	go build -o bin/server cmd/main.go

test:
	go test -v ./...

run: build
	./bin/server

migrate:
	docker run -v ./cmd/migrate/migrations/:/migrations migrate/migrate create -ext sql -dir /migrations $(word 2,$(MAKECMDGOALS))

migrate-up:
	go run cmd/migrate/migrate.go up

migrate-down:
	go run cmd/migrate/migrate.go down