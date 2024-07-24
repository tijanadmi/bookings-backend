DB_URL=postgresql://root:root@localhost:5433/bookings?sslmode=disable

postgres:
	docker run --name postgres16 -p 5433:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=root -d postgres:16-alpine

createdb:
	docker exec -it postgres16 createdb --username=root --owner=root bookings

dropdb:
	docker exec -it postgres16 dropdb bookings

migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

sqlc:
	sqlc generate

