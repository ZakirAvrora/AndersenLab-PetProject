postgres:
	docker run --name postgres15 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:15-alpine

createdb:
	docker exec -it postgres15 createdb --username=root --owner=root posts_api

dropdb:
	docker exec -it postgres15 dropdb posts_api

.PHONY: postgres createdb dropdb