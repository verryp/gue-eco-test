SHELL = /bin/bash

APP_NAME = monorepo-gue-eco
APP_DIR = $(shell pwd)

.PHONY: help
help:
	@echo 'Management commands for ${APP_NAME}:'
	@echo
	@echo 'Usage:'
	@echo '    make dep                                Update all dependencies'
	@echo '    make run                                Run all projects.'
	@echo '    make test                               Run tests on a compiled project.'
	@echo '    make compose.up                         Run all container and app.'
	@echo '    make compose.down                       Down all container.'
	@echo '    make compose.clean                      Clean all container.'
	@echo

.PHONY: dep
dep:
	@echo "Update all dependencies..."
	go get ./...

.PHONY: run
run: compose.up

.PHONY: compose.up
compose.up:
	@echo "Running all container..."
	docker-compose -f deployment/docker-compose.yaml --project-directory . up --build

.PHONY: compose.down
compose.down:
	@echo "down all container..."
	docker-compose -f deployment/docker-compose.yaml --project-directory . down

.PHONY: compose.clean
compose.clean:
	@echo "Clean all container..."
	docker-compose -f deployment/docker-compose.yaml --project-directory . stop
	docker-compose -f deployment/docker-compose.yaml --project-directory . rm
	docker-compose -f deployment/docker-compose.yaml --project-directory . down -v

test:
	@go test $$(go list ./... | grep -v /vendor/) -cover

build_plugin:
	@echo "Compiling plugins..."
	@go build -buildmode=plugin -v -o ./internal/gateway/plugins/authentication.so ./internal/gateway/plugins/authentication.go
	@echo "Done compiling plugin"