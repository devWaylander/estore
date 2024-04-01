FROM golang:1.22 as builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o bin/cart ./cmd/main.go

FROM debian

COPY --from=builder /app/bin/cart /cart

CMD [ "/cart" ]
