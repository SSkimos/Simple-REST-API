GO ?= go
DOCKER ?= docker

.PHONY: build

all: build up

build: reader_build

reader_build:
	$(MAKE) -C reader

up:
	docker-compose build && docker-compose up -d

down:
	docker-compose down --volumes

restart: down up