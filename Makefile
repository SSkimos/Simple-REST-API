GO ?= go
DOCKER ?= docker

.PHONY: build

all: build up

build: retrieval_build publisher_build reader_build cron_build

retrieval_build:
	$(MAKE) -C retrieval

publisher_build:
	$(MAKE) -C publisher

reader_build:
	$(MAKE) -C reader

cron_build:
	$(MAKE) -C cron

up:
	docker-compose build && docker-compose up -d

down:
	docker-compose down --volumes

restart: down build up