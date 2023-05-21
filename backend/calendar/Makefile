init:
	cp compose.override.yml.sample compose.override.yml

run:
	docker compose run --name=template_go_sqlc --rm gopher go run ./cmd

test:
	docker compose run --name=template_go_sqlc --rm gopher go test ./...

build:
	docker compose run --name=template_go_sqlc --rm gopher go build -o exec.exe .

fmt:
	docker compose run --name=template_go_sqlc --rm gopher go fmt ./...

sqlc:
	docker compose run --name=template_go_sqlc --rm gopher sqlc generate

migrate_file:
	migrate create -ext sql -dir migration migrate.sql

migrate_up:
	migrate -path migration -database "mysql://root:secret@tcp(localhost:3306)/aic" up

migrate_down:
	migrate -path migration -database "mysql://root:secret@tcp(localhost:3306)/aic" down


.PHONY: init run test build fmt sqlc migrate_file migrate_up migrate_down