init:
	sqlc init

generate:
	sqlc generate

table:
	docker run -it --rm --network host --volume "/Users/george/workspace/soleluxury-server/db:/db" migrate/migrate:v4.17.0 create -ext sql -dir /db/migrations $(name)

proto:
	rm -f pb/*.go
	protoc --proto_path=proto --go_out=pb --go_opt=paths=source_relative \
	--go-grpc_out=pb --go-grpc_opt=paths=source_relative \
	proto/*.proto

evans:
	docker run --rm -it -v "/Users/george/workspace/soleluxury-server:/mount:ro" \
    ghcr.io/ktr0731/evans:latest \
    --path /mount/proto/ \
    --proto soleluxury.proto \
    --host host.docker.internal \
    --port 8000 \
    repl

server:
	go run main.go

composeup:
	docker compose --env-file app.env up -d

composedown:
	docker compose down

PHONY: init generate table proto evans server composeup composedown