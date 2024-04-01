package main

import (
	"fmt"
	"log"
	"net/http"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	product "route256.ozon.ru/project/cart/external/product/client"
	stocks "route256.ozon.ru/project/cart/external/stocks/client"
	desc "route256.ozon.ru/project/cart/external/stocks/gen/api/orders/v1"
	"route256.ozon.ru/project/cart/internal/handler"
	"route256.ozon.ru/project/cart/internal/repo"
	"route256.ozon.ru/project/cart/internal/service"
	"route256.ozon.ru/project/cart/middleware"
)

const (
	ip               = "0.0.0.0"
	port             = "8080"
	grpcService      = "loms"
	grpcPort         = "50051"
	grpcClientHeader = "user"
)

func main() {
	// Repository
	repo := repo.New()

	// External Services
	productClient := product.New()

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", grpcService, grpcPort), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic(err)
	}
	client := desc.NewLOMSClient(conn)
	stocksClient := stocks.New(grpcClientHeader, client)

	// Service
	service := service.New(repo, productClient, stocksClient)

	// Handler
	mux := http.NewServeMux()
	handler.Configure(mux, service)
	wrappedMux := middleware.NewRequestLogger(mux)

	log.Printf("Server is up on: %s:%s", ip, port)
	if err := http.ListenAndServe(ip+":"+port, wrappedMux); err != nil {
		log.Fatal(err)
	}
}
