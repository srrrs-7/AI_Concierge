FROM golang:latest AS builder

WORKDIR /go/src

RUN go install -tags 'mysql' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

# ランナーステージ
FROM golang:latest AS runner

WORKDIR /go/src

COPY --from=builder /go/src .
