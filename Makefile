run-db:
	docker build -t balance_service_db build/postgres/
	docker run -dp 5432:5432 balance_service_db


migrate:
	migrate -path schema/ -database 'postgres://postgres:postgres@database:5432/balance_service?sslmode=disable' up

build:
	docker-compose build


run:
	docker-compose up
