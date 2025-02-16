db:
	docker run --name inkwave -p 5432:5432 -v ./db/data:/var/lib/postgresql/data -e POSTGRES_USER=test -e POSTGRES_PASSWORD=testsecret -d postgres:14-alpine

createdb:
	docker exec -it inkwave createdb --username test --owner=test inkwave

dropdb:
	docker exec -it inkwave dropdb inkwave

migrateup:
	migrate -path db/migration/ -database "postgresql://test:testsecret@localhost:5432/inkwave?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migration/ -database "postgresql://test:testsecret@localhost:5432/inkwave?sslmode=disable" -verbose down

new_migrate:
	@read -p "Enter migration name: " name; \
		migrate create -ext sql -dir ./db/migration -seq $$name

sqlc:
	sqlc generate -f ./config/sqlc.yml

test:
	go test -v -cover ./... 

		

.PHONY:  migrateup, migratedown, new_migrate, sqlc, createdb, dropdb, db
