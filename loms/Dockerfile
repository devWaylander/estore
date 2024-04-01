FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o bin/loms ./cmd/main.go

FROM debian

COPY --from=builder /app/bin/loms /loms

COPY --from=builder /app/stock-data.json ./stock-data.json

CMD [ "/loms" ]
