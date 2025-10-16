package aggregate

import (
	"log"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
)

type Aggregator struct {
	consumer *kafka.Consumer
}

func NewAggregator() (*Aggregator, error) {
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost",
		"group.id":          "counter-aggregator",
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		return nil, err
	}

	return &Aggregator{
		consumer: consumer,
	}, nil
}

func (agg *Aggregator) Run() error {
	log.Printf("Starting aggregator...")

	if err := agg.consumer.SubscribeTopics([]string{"counter"}, nil); err != nil {
		return err
	}

	for {
		msg, err := agg.consumer.ReadMessage(-1)
		if err == nil {
			log.Printf("Received message: %s\n", string(msg.Value))
		}
	}

}

func (agg *Aggregator) Close() {
	agg.consumer.Close()
}
