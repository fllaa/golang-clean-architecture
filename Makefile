.PHONY: clean critic security lint test build run watch swag

APP_NAME = golang_clean_architecture
BUILD_DIR = $(PWD)/build
MIGRATIONS_FOLDER = $(PWD)/db/migrations
# Set system environment variables
# DATABASE_URL = postgres://postgres:@localhost/golang_clean_architecture?sslmode=disable

clean:
	rm -rf ./build

critic:
	gocritic check -enableAll ./...

security:
	gosec ./...

lint:
	golangci-lint run ./...

test: clean critic security lint
	go test -v -timeout 30s -coverprofile=cover.out -cover ./...
	go tool cover -func=cover.out

build.web: test
	go build -ldflags="-w -s" -o $(BUILD_DIR)/web ./cmd/web/main.go

build.worker: test
	go build -ldflags="-w -s" -o $(BUILD_DIR)/worker ./cmd/worker/main.go

run.web: build.web
	$(BUILD_DIR)/web

run.worker: build.worker
	$(BUILD_DIR)/worker

watch.web: swag
	air -c .air-web.toml

watch.worker:
	air -c .air-worker.toml

migrate.up:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" up

migrate.down:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" down

migrate.force:
	migrate -path $(MIGRATIONS_FOLDER) -database "$(DATABASE_URL)" force $(version)

swag:
	swag init -g ./cmd/web/main.go