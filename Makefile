build-dev:
	docker build -t indexer -f containers/Dockerfile .

clean-dev:
	docker-compose -f containers/docker-compose.dev.yml down

run-dev:
	docker-compose -f containers/docker-compose.dev.yml up