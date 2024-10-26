SERVICE_NAME := app
COMPOSE := docker-compose
DOCKER_SYSTEM_PRUNE := docker system prune -f
TEST_SERVICE := test
DB_SERVICE := db
DB_TEST_SERVICE := db_test

.PHONY: run run-tests build up stop restart logs sh down cleanup help


run: 
	$(COMPOSE) up --build -d

run-tests: down
	$(COMPOSE) up --build $(TEST_SERVICE)

build:
	$(COMPOSE) build

up:
	$(COMPOSE) up -d $(SERVICE_NAME) $(DB_SERVICE)


stop:
	$(COMPOSE) stop

restart:
	$(COMPOSE) restart

logs:
	$(COMPOSE) logs -f $(SERVICE_NAME)

sh:
	$(COMPOSE) exec $(SERVICE_NAME) /bin/sh

down: 
	$(COMPOSE) down -v

cleanup: 
	$(DOCKER_SYSTEM_PRUNE)
