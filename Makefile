# Используем Bash как оболочку
SHELL := /bin/bash

include .env
export


# Запуск всех тестов
test:
	go test -tags mock,integration -race -cover ./...

# Запуск всех тестов с выключенным кешированием результата
test-no-cache:
	go test -tags mock,integration -race -cover -count=1 ./...

# Запуск всех линетров
lint:
	golangci-lint run ${args}

lint-fix:
	make lint args=--fix

redis-up:
	docker-compose up redis -d

redis-down:
	docker-compose down redis

redis-down-clean: redis-down
	@rm -rf db/redis

redis-full-restart: redis-down-clean redis-up

run-workers:
	go run main.go fetchers all

run-rest:
	go run main.go rest

# Генерация http-сервера на основе swagger-спецификации
# Требует предустановленного goswagger https://goswagger.io/
swagger:
	swagger generate server --exclude-main --exclude-spec -t internal/ -f api/swagger/file.yaml --name rest-server

install-sqlc:
	go install github.com/sqlc-dev/sqlc/cmd/sqlc@v1.21.0

install-golangci:
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.56.2
