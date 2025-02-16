db:
	docker run --name inkwave -p 5432:5432 -v ./db/data:/var/lib/postgresql/data -e POSTGRES_USER=test -e POSTGRES_PASSWORD=testsecret -d postgres:14-alpine

createdb:
	docker exec -it inkwave createdb --username test --owner=test inkwave

dropdb:
	docker exec -it inkwave dropdb inkwave
