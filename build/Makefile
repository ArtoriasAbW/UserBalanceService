run-db:
	docker build -t balance_service_db postgres/
	docker run -dp 5432:5432 balance_service_db


migrate:
	migrate -path ../schema -database 'postgres://admin:admin@localhost:5432/balance_service?sslmode=disable' up
