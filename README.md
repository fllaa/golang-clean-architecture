# Golang Clean Architecture Template

## Description

This is golang clean architecture template.

## Architecture

![Clean Architecture](architecture.png)

1. External system perform request (HTTP, gRPC, Messaging, etc)
2. The Delivery creates various Model from request data
3. The Delivery calls Use Case, and execute it using Model data
4. The Use Case create Entity data for the business logic
5. The Use Case calls Repository, and execute it using Entity data
6. The Repository use Entity data to perform database operation
7. The Repository perform database operation to the database
8. The Use Case create various Model for Gateway or from Entity data
9. The Use Case calls Gateway, and execute it using Model data
10. The Gateway using Model data to construct request to external system
11. The Gateway perform request to external system (HTTP, gRPC, Messaging, etc)

## Tech Stack

- Golang : <https://github.com/golang/go>
- PostgreSQL (Database) : <https://github.com/postgres/postgres>
- Apache Kafka : <https://github.com/apache/kafka>

## Framework & Library

- GoFiber (HTTP Framework) : <https://github.com/gofiber/fiber>
- GORM (ORM) : <https://github.com/go-gorm/gorm>
- Viper (Configuration) : <https://github.com/spf13/viper>
- Golang Migrate (Database Migration) : <https://github.com/golang-migrate/migrate>
- Go Playground Validator (Validation) : <https://github.com/go-playground/validator>
- Logrus (Logger) : <https://github.com/sirupsen/logrus>
- Confluent Kafka Golang : <https://github.com/confluentinc/confluent-kafka-go>

## Development Tools

- Air (Live Reload) : <https://github.com/air-verse/air>
- GoCritic : <https://github.com/go-critic/go-critic>
- GolangCI-Lint : <https://github.com/golangci/golangci-lint>
- GoSec : <https://github.com/securego/gosec>
- Migrate : <https://github.com/golang-migrate/migrate>
- Swaggo : <https://github.com/swaggo/swag>

## Configuration

All configuration is in `config.example.json` file. You can copy this file to `config.json` and modify it.

```shell
cp config.example.json config.json
```

## API Spec

All API Spec is in `api` folder.

## Database Migration

All database migration is in `db/migrations` folder.

### Create Migration

```shell
migrate create -ext sql -dir db/migrations create_table_xxx
```

### Run Migration

```shell
DATABASE_URL="postgres://postgres:5432@localhost/golang_clean_architecture?sslmode=disable" make migrate.up
```

## Run Application

### Run unit test

```bash
make test
```

### Run web server

```bash
make run.web
```

### Run worker

```bash
make run.worker
```
