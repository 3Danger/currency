# Используем Bash как оболочку
SHELL := /bin/bash

include .env
export


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

