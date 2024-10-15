FROM golang:1.22.6-alpine AS base
FROM base AS dev

RUN apk add --no-cache git bash

WORKDIR /app

COPY . .

RUN go mod download


RUN go build -o ./bin/api ./cmd/main.go

CMD ["./bin/api"]
