coverage:
	go test -coverprofile=coverage.out ./...
	gocover-cobertura < coverage.out > coverage.xml
format-sql:
	@for file in db/query/*.sql; do \
		pg_format "$$file" -o "$$file"; \
	done
postgres:
	docker run --name postgres12 -p 5432:5432 \
		-e POSTGRES_USER=myuser \
		-e POSTGRES_PASSWORD=mypassword \
		-d postgres:12-alpine

createdb:
	docker exec -it postgres12 createdb --username=myuser --owner=myuser simple_bank

dropdb:
	docker exec -it postgres12 dropdb simple_bank

migrateup:
	migrate -path db/migration -database "postgresql://myuser:mypassword@localhost:5432/simple_bank?sslmode=disable" --verbose up

migratedown:
	migrate -path db/migration -database "postgresql://myuser:mypassword@localhost:5432/simple_bank?sslmode=disable" --verbose down

sqlc:
	sqlc generate

test:
	go test -v -cover ./...

server:
	go run main.go

mock: 
	mockgen -package mockdb -destination db/mock/store.go github.com/development/simple-bank/db/sqlc Store
	
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test server mock
