postgres:
	docker run --name postgres-12 --network grocery-network -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

createdb:
	docker exec -it postgres-12 createdb --username=root --owner=root good_grocer

dropdb:
	docker exec -it postgres-12 dropdb good_grocer

server:
	go run main.go