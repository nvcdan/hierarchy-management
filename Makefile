SERVICE_NAME := app 

run:
	docker-compose up --build -d

run-tests:
	docker-compose run $(SERVICE_NAME) go test -v ./...

build:
	docker-compose build

up:
	docker-compose up -d

stop:
	docker-compose stop

restart:
	docker-compose restart

logs:
	docker-compose logs -f $(SERVICE_NAME)

sh:
	docker-compose exec $(SERVICE_NAME) /bin/sh

down:
	docker-compose down -v

cleanup:
	docker system prune -f

help:
	@echo "Makefile commands for Docker Compose:"
	@echo "  build     - Build the Docker images"
	@echo "  up        - Start services in detached mode"
	@echo "  stop      - Stop the services"
	@echo "  restart   - Restart the services"
	@echo "  logs      - View logs of the services"
	@echo "  sh        - Access the service's shell"
	@echo "  down      - Stop and remove services, networks, and volumes"
	@echo "  cleanup   - Remove unused Docker images and containers"