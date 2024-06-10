# Используем Bash как оболочку
SHELL := /bin/bash

include .env
export


go-dependices:
	go mod tidy

setup: go-dependices postgres-up migrate-postgres-up

migrate-postgres-up:
	go run main.go migrate postgres up

postgres-up:
	docker-compose up postgres -d

postgres-down:
	docker-compose down postgres

postgres-down-clean: postgres-down
	@rm -rf db/postgresql/data

postgres-full-restart: postgres-down-clean postgres-up

redis-up:
	docker-compose up redis -d

redis-down:
	docker-compose down redis

redis-down-clean: redis-down
	@rm -rf ./db/redis

redis-full-restart: redis-down-clean redis-up

run-workers:
	go run main.go fetchers all

run-rest:
	go run main.go rest

update-swagger:
	swag init --output api/swagger

install-swag:
	go install github.com/swaggo/swag/cmd/swag@latest

