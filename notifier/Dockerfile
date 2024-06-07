FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o bin/notifier ./cmd

FROM debian

COPY --from=builder /app/bin/notifier /notifier

# COPY --from=builder /app/.env ./.env

ENTRYPOINT ["/notifier", "-bootstrap-server=kafka0:29092"]
