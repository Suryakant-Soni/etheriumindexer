.PHONY: build run all

build-dev:
	docker build -t indexer -f containers/Dockerfile.dev .

clean-dev:
	docker-compose -f containers/docker-compose.dev.yml down

run-dev:
	docker-compose -f containers/docker-compose.dev.yml up

all: 
	build-dev run-dev