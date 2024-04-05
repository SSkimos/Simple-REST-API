GO ?= go
DOCKER ?= docker

.PHONY: build

build:
	$(GO) build -o ./build/main  ./cmd/main.go
#
#up:
#	docker-compose build && docker-compose up -d
#
#down:
#	docker-compose down --volumes
#
#restart: down up