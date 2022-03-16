psql:
	sudo docker run -d -p 6565:5432 -e POSTGRES_PASSWORD=postgres -e POSTGRES_USER=postgres --name psql_todolist postgres:14-alpine3.15

rm_psql:
	sudo docker rm -f psql_todolist

createdb:
	sudo docker exec -it psql_todolist createdb -U postgres simple_bank

dropdb:
	sudo docker exec -it psql_todolist dropdb -U postgres simple_bank

opendb:
	sudo docker exec -it psql_todolist /bin/bash

migrateup:
	sudo migrate -path ~/go/src/github/todolist/schema -database postgresql://postgres:postgres@localhost:6565/simple_bank?sslmode=disable -verbose up

migratedown:
	sudo migrate -path ~/go/src/github/todolist/schema -database postgresql://postgres:postgres@localhost:6565/simple_bank?sslmode=disable -verbose down

PHONY: psql, rm_psql, createdb, dropdb, opendb, migrateup, migratedown
