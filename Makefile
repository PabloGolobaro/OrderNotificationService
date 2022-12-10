BIN := blueprint

# Where to push the docker image.
REGISTRY ?= https://hub.docker.com/repositories

# This version-strategy uses git tags to set the version string
#VERSION := $(shell git describe --tags --always --dirty)

# This version-strategy uses a manual value to set the version string
include .env

prod:
	docker-compose up

prod_build:
	docker-compose build

server:
	docker-compose up --no-start --force-recreate server
	docker push pablogolobar/order_server:$(VERSION)
