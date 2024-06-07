package main

import (
	"route256.ozon.ru/project/notifier/infra/kafka"
	"route256.ozon.ru/project/notifier/infra/kafka/consumer_group"
)

type config struct {
	kafka         kafka.Config
	consumerGroup consumer_group.Config
}

func newConfig(f flags) config {
	return config{
		kafka: kafka.Config{
			Brokers: []string{
				f.bootstrapServer,
			},
		},
		consumerGroup: consumer_group.Config{
			GroupName: f.consumerGroupName,
			Topics:    []string{f.topic},
		},
	}
}
