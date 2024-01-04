createdb:
	createdb --username=postgres --owner=postgres go_finance

postgres:
	docker run --name postgres -p 5433:5432 -e POSTGRES_PASSWORD=postgres -d postgres:14-alpine

migrateup:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5433/go_finance?sslmode=disable" -verbose up

migrationdrop:
	migrate -path db/migration -database "postgresql://postgres:postgres@localhost:5433/go_finance?sslmode=disable" -verbose down

test:
	go test -v -cover ./...

server:
	go run main.go

sqlc-gen:
	docker run --rm -v $$(pwd):/src -w /src kjconroy/sqlc generate

server:
	go run main.go

.PHONY: createdb postgres dropdb migrateup migrationdrop test server sqlc-gen server