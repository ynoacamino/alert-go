FROM golang:latest AS builder

WORKDIR /app

COPY ./go.mod /app/go.mod
COPY ./go.sum /app/go.sum

RUN go install github.com/pressly/goose/v3/cmd/goose@latest

RUN go mod download

COPY . /app

RUN goose -dir /app/db/schemas sqlite3 /app/db/database.db up

RUN go build -o main /app/cmd/main.go

FROM debian:latest AS production

WORKDIR /app

COPY --from=builder /app/db/database.db /app/db/database.db
COPY --from=builder /app/main /app/main

CMD ["/app/main"]
