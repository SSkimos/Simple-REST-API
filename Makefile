GO ?= go
DOCKER ?= docker

.PHONY: build

all: build up

build: reader_build publisher_build

reader_build:
	$(MAKE) -C reader

publisher_build:
	$(MAKE) -C publisher

up:
	docker-compose build && docker-compose up -d

down:
	docker-compose down --volumes

restart: down up