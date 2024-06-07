package kafka

import (
	"flag"
)

type Flags struct {
	startID         int
	bootstrapServer string
	topic           string
}

var CliFlags = Flags{}

func init() {
	flag.IntVar(&CliFlags.startID, "start-id", 1, "start order-id of all messages")
	flag.StringVar(&CliFlags.bootstrapServer, "bootstrap-server", "kafka0:29092", "kafka broker host and port")
	flag.StringVar(&CliFlags.topic, "topic", "loms.order-events", "topic to produce")

	flag.Parse()
}
