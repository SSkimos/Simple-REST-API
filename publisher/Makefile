GO ?= go
DOCKER ?= docker
BUILD_DIR=$(CURDIR)/build

.PHONY: build
build:
	$(GO) build -o $(BUILD_DIR)/main  ./cmd/main.go

up:
	docker-compose build && docker-compose up -d

down:
	docker-compose down