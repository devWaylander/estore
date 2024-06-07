package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/IBM/sarama"
	"route256.ozon.ru/project/notifier/infra/kafka/consumer_group"
	"route256.ozon.ru/project/notifier/util"
)

func main() {
	// Graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)

		<-c
		cancel()
	}()

	var (
		wg   = &sync.WaitGroup{}
		conf = newConfig(cliFlags)
	)

	fmt.Printf("%+v\n", conf)
	cg, err := consumer_group.NewConsumerGroup(
		conf.kafka,
		conf.consumerGroup,
		consumer_group.NewConsumerGroupHandler(),
		consumer_group.WithOffsetsInitial(sarama.OffsetOldest),
	)
	if err != nil {
		log.Fatal(err)
	}

	cg.Run(ctx, wg)
	wg.Wait()

	// Graceful shutdown
	g, gCtx := util.EGWithContext(ctx)
	g.Go(func() error {
		<-gCtx.Done()
		log.Printf("Sync sarama producer is shutting down")
		cg.Close()
		return err
	})

	if err := g.Wait(); err != nil {
		log.Printf("exit reason: %s \\n", err)
	}
}
