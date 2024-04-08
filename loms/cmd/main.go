package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pgx "github.com/jackc/pgx/v5"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	desc "route256.ozon.ru/project/loms/gen/api/orders/v1"
	controller "route256.ozon.ru/project/loms/internal/app"
	orders_repo "route256.ozon.ru/project/loms/internal/repo/db_repo/orders"
	stocks_repo "route256.ozon.ru/project/loms/internal/repo/db_repo/stocks"
	"route256.ozon.ru/project/loms/internal/service"
	"route256.ozon.ru/project/loms/middleware"
)

const (
	grpcPort       = "50051"
	jsonStockPath  = "../stock-data.json"
	ordersDBString = "postgres://postgres:12341234@127.0.0.1:1234/orders?sslmode=disable"
	stocksDBString = "postgres://postgres:12341234@127.0.0.1:1234/stocks?sslmode=disable"
)

// func headerMatcher(key string) (string, bool) {
// 	switch strings.ToLower(key) {
// 	case "x-auth":
// 		return key, true
// 	default:
// 		return key, false
// 	}
// }

func main() {
	ctx := context.Background()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", grpcPort))
	if err != nil {
		panic(err)
	}
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middleware.Panic,
			middleware.Logger,
			middleware.Validate,
		),
	)
	reflection.Register(grpcServer)

	// jsonStock, err := os.Open(jsonStockPath)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// defer jsonStock.Close()
	// memoryRepo := memory_repo.New(jsonStock)

	dbStocksConn, err := pgx.Connect(ctx, stocksDBString)
	if err != nil {
		panic(err)
	}
	defer dbStocksConn.Close(ctx)
	repoStocks := stocks_repo.New(dbStocksConn)

	dbOrdersConn, err := pgx.Connect(ctx, ordersDBString)
	if err != nil {
		panic(err)
	}
	defer dbOrdersConn.Close(ctx)
	repoOrders := orders_repo.New(dbOrdersConn)

	service := service.New(repoOrders, repoStocks)
	controller := controller.New(service)

	desc.RegisterLOMSServer(grpcServer, controller)

	log.Printf("Server listening at %v", lis.Addr())
	if err = grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	// go func() {
	// 	if err = grpcServer.Serve(lis); err != nil {
	// 		log.Fatalf("failed to serve: %v", err)
	// 	}
	// }()

	// conn, err := grpc.Dial(fmt.Sprintf(":%s", grpcPort), grpc.WithBlock(), grpc.WithTransportCredentials(insecure.NewCredentials()))
	// if err != nil {
	// 	log.Fatalln("Failed to deal: ", err)
	// }

	// gwmux := runtime.NewServeMux()

	// gwmux.Handle("/swaggerui", )
	// gwmux = runtime.NewServeMux(runtime.WithIncomingHeaderMatcher(headerMatcher))
	// if err = desc.RegisterLOMSHandler(context.Background(), gwmux, conn); err != nil {
	// 	log.Fatalln("Failed to register gateway: ", err)
	// }
	// gwServer := &http.Server{
	// 	Addr:    fmt.Sprintf(":%s", httpPort),
	// 	Handler: middleware.WithHTTPLoggingMiddleware(gwmux),
	// }

	// log.Printf("Serving gRPC-Gateway on %s\n", gwServer.Addr)
	// log.Fatalln(gwServer.ListenAndServe())
}
