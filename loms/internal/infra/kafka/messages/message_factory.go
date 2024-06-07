package messages

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/IBM/sarama"
	"route256.ozon.ru/project/loms/internal/infra/kafka"
	"route256.ozon.ru/project/loms/internal/infra/kafka/events"
)

type operationMomentGenerator interface {
	Generate() time.Time
}

type MessageFactory struct {
	OperationMoment operationMomentGenerator
}

func (f *MessageFactory) Create(event events.Event, config kafka.Config) *sarama.ProducerMessage {
	bytes, err := json.Marshal(event)
	if err != nil {
		log.Fatal(err)
	}

	return &sarama.ProducerMessage{
		Topic: config.Producer.Topic,
		Key:   sarama.StringEncoder(strconv.FormatInt(event.ID, 10)),
		Value: sarama.ByteEncoder(bytes),
		Headers: []sarama.RecordHeader{
			{
				Key:   []byte("loms"),
				Value: []byte("loms-sync-prod"),
			},
		},
		Partition: 1,
		Timestamp: time.Now(),
	}
}

func NewDefaultFactory() *MessageFactory {
	return New(
		&Clock{},
	)
}

func New(
	momentGen operationMomentGenerator,
) *MessageFactory {
	return &MessageFactory{
		momentGen,
	}
}

type Clock struct{}

func (c *Clock) Generate() time.Time {
	return time.Now()
}
