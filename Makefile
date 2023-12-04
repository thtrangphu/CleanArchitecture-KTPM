postgresinit:
	docker run --name postgres15 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:15-alpine

postgres:
# Open pg cmd
#  /l: all table
	docker exec -it postgres15 psql

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root clean-architecture

dropdb:
	docker exec -it postgres15 dropdb clean-architecture

migrateup:
	migrate  -path .\database\migrations\ -datebase "postgresql://root:password@localhost:5433/clean-architecture?sslmode=disable" -verbose up
	
migratedown:
	migrate  -path .\database\migrations\ -datebase "postgresql://root:password@localhost:5433/clean-architecture?sslmode=disable" -verbose down

.PHONY: postgresinit postgres createdb dropdb migrateup migratedown