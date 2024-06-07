package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/IBM/sarama"
	pgx "github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	desc "route256.ozon.ru/project/loms/gen/api/orders/v1"
	controller "route256.ozon.ru/project/loms/internal/app"
	"route256.ozon.ru/project/loms/internal/infra/kafka"
	"route256.ozon.ru/project/loms/internal/infra/kafka/producer"
	orders_repo "route256.ozon.ru/project/loms/internal/repo/db_repo/orders"
	stocks_repo "route256.ozon.ru/project/loms/internal/repo/db_repo/stocks"
	"route256.ozon.ru/project/loms/internal/service"
	"route256.ozon.ru/project/loms/middleware"
	"route256.ozon.ru/project/loms/util"
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
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("failed to load .env file")
	}

	// Graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		c := make(chan os.Signal, 2)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		<-c
		cancel()
	}()

	// GRPC server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", os.Getenv("PORT")))
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

	// in-memory repo
	// jsonStock, err := os.Open(jsonStockPath)
	// if err != nil {
	// 	log.Panic(err)
	// }
	// defer jsonStock.Close()
	// memoryRepo := memory_repo.New(jsonStock)

	// Postgres
	dbStocksConn, err := pgx.Connect(ctx, os.Getenv("STOCKS_DB_STRING"))
	if err != nil {
		panic(err)
	}
	defer dbStocksConn.Close(ctx)
	repoStocks := stocks_repo.New(dbStocksConn)

	dbOrdersConn, err := pgx.Connect(ctx, os.Getenv("ORDERS_DB_STRING"))
	if err != nil {
		panic(err)
	}
	defer dbOrdersConn.Close(ctx)
	repoOrders := orders_repo.New(dbOrdersConn)

	// Kafka
	conf := kafka.NewConfig(kafka.CliFlags)

	prod, err := producer.NewSyncProducer(conf.Kafka,
		producer.WithIdempotent(),
		producer.WithRequiredAcks(sarama.WaitForAll),
		producer.WithMaxOpenRequests(1),
		producer.WithMaxRetries(5),
		producer.WithRetryBackoff(10*time.Millisecond),
		//producer.WithProducerPartitioner(sarama.NewManualPartitioner),
		//producer.WithProducerPartitioner(sarama.NewRoundRobinPartitioner),
		producer.WithProducerPartitioner(sarama.NewRandomPartitioner),
	)
	if err != nil {
		log.Fatal(err)
	}

	// App run
	service := service.New(repoOrders, repoStocks, prod, conf)
	controller := controller.New(service)
	desc.RegisterLOMSServer(grpcServer, controller)

	// Graceful shutdown
	g, gCtx := util.EGWithContext(ctx)
	g.Go(func() error {
		log.Printf("GRPC server listening at %v", lis.Addr())
		return grpcServer.Serve(lis)
	})
	g.Go(func() error {
		<-gCtx.Done()
		log.Printf("Sync sarama producer is shutting down")
		prod.Close()
		return err
	})
	g.Go(func() error {
		<-gCtx.Done()
		log.Printf("GRPC server is shutting down")
		grpcServer.GracefulStop()
		return err
	})

	if err := g.Wait(); err != nil {
		log.Printf("exit reason: %s \\n", err)
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
