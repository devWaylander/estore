package main

import (
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	desc "route256.ozon.ru/project/loms/gen/api/orders/v1"
	controller "route256.ozon.ru/project/loms/internal/app"
	"route256.ozon.ru/project/loms/internal/repo"
	"route256.ozon.ru/project/loms/internal/service"
	"route256.ozon.ru/project/loms/middleware"
)

const (
	grpcPort      = "50051"
	jsonStockPath = "../stock-data.json"
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

	jsonStock, err := os.Open(jsonStockPath)
	if err != nil {
		log.Panic(err)
	}
	defer jsonStock.Close()

	repo := repo.New(jsonStock)
	service := service.New(repo)
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
