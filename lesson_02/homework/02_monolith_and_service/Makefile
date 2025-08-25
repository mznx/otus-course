build:
	docker build ./monolith/app -t monolith_02_02:latest
	docker build ./dialog-service/app -t dialog-service_02_02:latest

start:
	docker compose -f ./docker-compose-dev.yml up -d

stop:
	docker compose -f ./docker-compose-dev.yml down
