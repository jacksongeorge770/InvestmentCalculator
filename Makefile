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