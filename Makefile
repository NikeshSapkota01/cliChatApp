docker/image:
	docker build . -t go_cli_client:1.0

docker/run:
	docker run go_cli_client:1.0

docker/stop:
	docker compose down

docker/start:
	docker compose up 

postgres:
	docker exec -it goCliChat psql -U postgres

createdb:
	docker exec -it goCliChat createdb --username=postgres --owner=postgres chatApi

dropdb:
	docker exec -it goCliChat dropdb chatApi

migrateup:
	migrate -path db/migrations -database "postgresql://postgres:root@localhost:5440/chatApi?sslmode=disable" -verbose up

migratedown:
	migrate -path db/migrations -database "postgresql://postgres:root@localhost:5433/chatApi?sslmode=disable" -verbose down

.PHONY: docker/image docker/run docker/stop docker/start postgres createdb dropdb migrateup migratedown