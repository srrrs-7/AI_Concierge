
db_url=mysql://root:secret@tcp(auth_db:3306)/aic


init:
	cp compose.override.yml.sample compose.override.yml

download:
	docker compose run --name=auth --rm auth_api go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
	docker compose run --name=auth --rm auth_api go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

run:
	docker compose run --name=auth --rm auth_api go run ./cmd

build:
	docker compose run --name=auth --rm auth_api go build -o /go/src/auth.exe ./cmd

exec:
	docker compose run --name=auth --rm auth_api /go/src/auth.exe

test:
	docker compose run --name=auth --rm auth_api go test ./...

fmt:
	docker compose run --name=auth --rm auth_api go fmt ./...

sqlc:
	docker compose run --name=auth --rm auth_api sqlc generate

migrate_file:
	docker compose run --name=auth --rm auth_api migrate create -ext sql -dir _migration migrate

migrate_up:
	docker compose run --name=auth --rm auth_api migrate -path _migration -database "$(db_url)" up

migrate_down:
	docker compose run --name=auth --rm auth_api migrate -path _migration -database "$(db_url)" down

bucket:
	docker compose run --name=auth --rm aws s3 mb s3://aic --endpoint-url=http://localstack:4566


.PHONY: init run test build fmt sqlc migrate_file migrate_up migrate_down