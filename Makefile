.PHONY: build run-dev run-prod

build:
	docker-compose build

run-dev:
	docker-compose up app_dev


run-prod:
	docker-compose up app_prod
