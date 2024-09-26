FROM golang:1.22.6-alpine AS base
FROM base AS dev

RUN apk add --no-cache git

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o ./bin/api ./cmd/api/main.go

FROM alpine:latest AS prod

WORKDIR /app
COPY --from=dev /app/api .

RUN adduser -D -g '' appuser && chown -R appuser /app
USER appuser

CMD ["./api"]
