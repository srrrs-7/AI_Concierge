FROM golang:latest

WORKDIR /go/src

COPY . .

RUN go mod download

RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

RUN go build -o /go/bin/auth.exe ./cmd

CMD go run ./cmd
