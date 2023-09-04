SHELL = /bin/bash

APP_NAME = monorepo-gue-eco
APP_DIR = $(shell pwd)

.PHONY: help
help:
	@echo 'Management commands for ${APP_NAME}:'
	@echo
	@echo 'Usage:'
	@echo '    make build                              Compile the project.'
	@echo '    make run ARGS=                          Run with supplied arguments.'
	@echo '    make test                               Run tests on a compiled project.'
	@echo '    make compose.up                         Run all container and app.'
	@echo '    make compose.down                       Down all container.'
	@echo '    make compose.clean                      Clean all container.'
	@echo

.PHONY: dep
dep:
	@echo "Update all dependencies..."
	go get ./...

.PHONY: build
build: dep
	@echo "Building ${APP_DIR}"
	go build -ldflags "-w" -o bin/${APP_NAME}/product ${APP_DIR}/cmd/product/main.go

.PHONY: run
run:
	@echo "Running ${APP_NAME}..."
	go run ${APP_DIR}/main.go ${ARGS}

.PHONY: compose.up
compose.up:
	@echo "Running all container..."
	docker-compose -f deployment/docker-compose.yaml --project-directory . up -d --build

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