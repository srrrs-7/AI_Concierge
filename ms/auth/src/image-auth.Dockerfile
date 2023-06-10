# ビルダーステージ
FROM golang:latest AS builder

WORKDIR /go/src

COPY . .

RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest
RUN go install github.com/kyleconroy/sqlc/cmd/sqlc@latest

RUN go mod download

# ランナーステージ
FROM golang:latest AS runner

WORKDIR /go/src

COPY --from=builder /go/src .

CMD [ "go", "run", "./cmd" ]
