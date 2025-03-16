postgres:
	docker run --name myEmail3 -p 5435:5432 -e POSTGRES_USER=admin -e POSTGRES_PASSWORD=aceanh258 -d postgres:12-alpine

createdb:
	docker exec -it myEmail1 createdb --username=admin --owner=admin newPostgres

dropdb:
	docker exec -it myEmail1 dropdb newPostgres

migrateup:
	migrate -path db/postGresMigration -database "postgres://admin:aceanh258@localhost:5433/postgres?sslmode=disable" -verbose up

migratedown:
	migrate -path db/postGresMigration -database "postgres://admin:aceanh258@localhost:5433/postgres?sslmode=disable" -verbose down

test:
	go test -v -cover ./...
	
.PHONY: postgres createdb dropdb migrateup migratedown sqlc test
