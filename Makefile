docker/image:
	docker build -t go_cli_client:1.0

docker/run:
	docker run go_cli_client:1.0

docker/stop:
	docker compose down

docker/start:
	docker compose up --build --remove-orphans

postgres:
	docker exec -it goCliChat psql -U postgres

createdb:
	docker exec -it goCliChat createdb --username=postgres --owner=root chatApi

dropdb:
	docker exec -it goCliChat dropdb chatApi

.PHONY: docker/image docker/run docker/stop docker/start postgres createdb dropdb