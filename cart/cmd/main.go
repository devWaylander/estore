package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	product "route256.ozon.ru/project/cart/external/product/client"
	stocks "route256.ozon.ru/project/cart/external/stocks/client"
	desc "route256.ozon.ru/project/cart/external/stocks/gen/api/orders/v1"
	"route256.ozon.ru/project/cart/internal/handler"
	"route256.ozon.ru/project/cart/internal/repo"
	"route256.ozon.ru/project/cart/internal/service"
	"route256.ozon.ru/project/cart/middleware"
	"route256.ozon.ru/project/cart/util"
)

func main() {
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	var (
		ip   = os.Getenv("IP")
		port = os.Getenv("PORT")
	)

	// Graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		<-c
		cancel()
	}()

	// Repository
	repo := repo.New()

	// External Services
	productClient := product.New()

	conn, err := grpc.Dial(fmt.Sprintf("%s:%s", os.Getenv("LOMS_SERVICE"), os.Getenv("LOMS_PORT")), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic(err)
	}
	client := desc.NewLOMSClient(conn)
	stocksClient := stocks.New(os.Getenv("GRPC_CLIENT_HEADER"), client)

	// Service
	service := service.New(repo, productClient, stocksClient)

	// Handler
	mux := http.NewServeMux()
	handler.Configure(ctx, mux, service)
	wrappedMux := middleware.NewRequestLogger(mux)

	httpServer := &http.Server{
		Addr:    fmt.Sprintf(":%s", port),
		Handler: wrappedMux,
	}

	g, gCtx := util.EGWithContext(ctx)
	g.Go(func() error {
		log.Printf("Server is up on: %s:%s", ip, port)
		return httpServer.ListenAndServe()
	})
	g.Go(func() error {
		<-gCtx.Done()
		log.Printf("Server is shutting down: %s:%s", ip, port)
		return httpServer.Shutdown(context.Background())
	})

	if _, err := g.Wait(); err != nil {
		log.Printf("exit reason: %s \\n", err)
	}
}
