build:
	@mkdir -p bin
	@go build -o bin/cal cmd/main.go

build-migrate:
	@mkdir -p bin
	@go build -o bin/migrate cmd/migrate/main.go

test:
	@go test -v ./...

run: build
	@./bin/cal

migrate-up:
	@echo "Running migration up..."
	@go run cmd/migrate/main.go up

migrate-down:
	@echo "Running migration down..."
	@go run cmd/migrate/main.go down

migrate:
	@echo "Running migration..."
	@goose -dir cmd/migrate/migrations mysql://${DB_USER}:${DB_PASSWORD}@tcp(${DB_HOST}:${DB_PORT})/${DB_NAME}?parseTime=true up