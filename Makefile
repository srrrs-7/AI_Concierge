# initialize
init:
	cp compose.override.yaml.sample compose.override.yaml
up:
	docker compose up -d --build
down:
	docker compose down
.PHONY: init up down

# UI
ui: npm-install ui-start
npm-install:
	docker compose run --rm ui npm install
ui-start:
	docker compose up ui storybook -d --build
ui-sb-test:
	npm run test-storybook
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
db_url="mysql://root:secret@tcp(auth-db:3306)/aic"

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
	docker compose run --rm auth-api go fmt ./...
auth-migrate-file:
	docker compose run --rm auth-api migrate create -ext sql -dir _migration migrate
auth-migrate-up:
	docker compose run --rm auth-api migrate -path _migration -database $(db_url) up
auth-migrate-down:
	docker compose run --rm auth-api migrate -path _migration -database $(db_url) down
.PHONY: auth auth-build auth-exec auth-test auth-coverage auth-fmt auth-sqlc auth-migrate-file auth-migrate-up auth-migrate-down

# localstack
localstack_URL="http://localstack:4566"
localstack:
	localstack status services
.PHONY: localstack

# S3
bucket="aic"
file="README.md"
bucket:
	docker compose run --rm aws s3 mb s3://$(bucket) --endpoint-url=$(localstack_URL)
upload:
	docker compose run --rm aws s3 cp $(file) s3://my-bucket/$(bucket)
download:
	docker compose run --rm aws s3 cp s3://$(bucket)/$(file) ./
.PHONY: bucket upload download

# SQS
queue_URL = "http://localstack:4566/000000000000/aic"
message="Hi!! SQS!"
sqs:
	docker compose run --rm aws sqs create-queue --queue-name aic --endpoint-url=$(localstack_URL)
send:
	docker compose run --rm aws sqs send-message --queue-url=$(queue_URL) --message-body=$(message) --endpoint-url=$(localstack_URL)
receive:
	docker compose run --rm aws sqs receive-message --queue-url=$(queue_URL) --endpoint-url=$(localstack_URL)
.PHONY: sqs sqs-send sqs-receive

# PlantUML
plant-uml:
	docker compose up plant-uml -d --build
.PHONY: uml

# Redoc
redoc:
	docker compose up redoc -d --build
.PHONY: redoc
