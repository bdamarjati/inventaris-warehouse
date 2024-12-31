
include .env

createdb:
	sqlite3 $(DB_NAME) "VACUUM;"
dropdb:
	del $(DB_NAME)
migrateup:
	migrate -database "sqlite3://$(DB_NAME)" -path "db/migrations" -verbose up
migratedown:
	migrate -database "sqlite3://$(DB_NAME)" -path "db/migrations" -verbose down
test:
	go test -v -cover ./...

.PHONY: createdb dropdb migrateup migratedown test
