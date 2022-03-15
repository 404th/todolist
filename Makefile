

# postgres:
# 	sudo docker run --name pgs -p 5555:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=pass -d postgres:14-alpine3.15

# createdb:
# 	sudo docker exec -it pgs createdb --username=postgres --owner=postgres simple_bank

# dropdb    :
# 	sudo docker exec -it pgs dropdb -U postgres simple_bank

# migrateup:
# 	sudo migrate -path ./schema -database "postgresql://postgres:pass@localhost:5555/simple_bank?sslmode=disable" -verbose up

# # static: ./schema ( migration path ) | postgresql ( db driver )
# # pgs ( docker container ) | postgres ( db username ) | simple_bank ( db name )	
# migratedown:
# 	sudo migrate -path ./schema -database "postgresql://postgres:pass@localhost:5555/simple_bank?sslmode=disable" -verbose down

# # static: pgs ( docker container ) | postgres ( db username ) | simple_bank ( db name )
# opendb:
# 	sudo docker exec -it pgs psql -U postgres simple_bank

# .PHONY: createdb, dropdbm, migrateup, migratedown, opendb

postgres:
	sudo docker run -d --name pgs -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres postgres:14-alpine3.15

createdb:
	sudo docker exec -it pgs createdb --username=postgres --owner=postgres simple_bank

dropdb:
	sudo docker exec -it dropdb -U postgres simple_bank

migrateup:
	sudo migrate -path ./migration -database "postgresql://postgres:postgres@localhost:5555/simple_bank?sslmode=disable" -verbose up

migratedown:
	sudo migrate -path ./migration -database "postgresql://postgres:postgres@localhost:5555/simple_bank?sslmode=disable" -verbose down

opendb:
	sudo docker exec -it pgs psql -U postgres simple_bank

.PHONY: postgres, createdb, dropdb, migrateup, migratedown, opendb
