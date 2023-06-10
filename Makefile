# initialize
init:
	cp compose.override.yaml.sample compose.override.yaml
up:
	docker compose up -d --build
down:
	docker compose down
.PHONY:init up down


# UI
ui:
	docker compose up ui storybook -d --build
ui-sb-test:
	docker compose run --rm storybook npx playwright install
	docker compose run --rm storybook npm run test-storybook
ui-fmt:
	docker compose run --rm ui npx prettier --write ./src
ui-lint:
	docker compose run --rm ui npx eslint --fix ./src
ui-test:
	docker compose run --rm ui npm run test
ui-build:
	docker compose run --rm ui npm run build
.PHONY: ui ui-dev ui-sb ui-sb-test ui-fmt ui-lint ui-test ui-build


# Auth
db_url=mysql://root:secret@tcp(auth_db:3306)/aic

auth:
	docker compose up auth-api auth-db auth-cache localstack aws -d --build
auth-build:
	docker compose run --rm auth-api go build -o /go/bin/auth.exe ./cmd
auth-exec:
	docker compose run --rm auth-api /go/bin/auth.exe
auth-test:
	docker compose run --rm auth-api go test -cover ./... -coverprofile=/go/log/cover.out
auth-coverage:
	docker compose run --rm auth-api go tool cover -html=/go/log/cover.out -o /go/log/cover.html
	open ./ms/auth/log/cover.html
auth-fmt:
	docker compose run --name=auth --rm auth-api go fmt ./...
auth-sqlc:
	docker compose run --name=auth --rm auth-api sqlc generate
auht-migrate-file:
	docker compose run --name=auth --rm auth-api migrate create -ext sql -dir _migration migrate
auht-migrate-up:
	docker compose run --name=auth --rm auth-api migrate -path _migration -database "$(db_url)" up
auht-migrate-down:
	docker compose run --name=auth --rm auth-api migrate -path _migration -database "$(db_url)" down
auht-bucket:
	docker compose run --name=auth --rm aws s3 mb s3://aic --endpoint-url=http://localstack:4566
.PHONY: auth auth-build auth-exec auth-test auth-coverage auth-fmt auth-sqlc auht-migrate-file auht-migrate-up auht-migrate-down


# PlantUML
uml:
	docker compose up plant-uml -d --build
.PHONY: uml