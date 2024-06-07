package kafka

type appConfig struct {
	// repeatCnt int
	StartID int
	// count     int
}

type producerConfig struct {
	Topic string
}

type KafkaConfig struct {
	Brokers []string
}

type Config struct {
	App      appConfig
	Kafka    KafkaConfig
	Producer producerConfig
}

func NewConfig(f Flags) Config {
	return Config{
		App: appConfig{
			StartID: f.startID,
		},
		Kafka: KafkaConfig{
			Brokers: []string{
				f.bootstrapServer,
			},
		},
		Producer: producerConfig{
			Topic: f.topic,
		},
	}
}
