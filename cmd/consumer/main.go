package main

import (
	"context"
	"fmt"

	"github.com/IBM/sarama"
)

type handler struct{}

func (h *handler) Setup(sarama.ConsumerGroupSession) error   { return nil }
func (h *handler) Cleanup(sarama.ConsumerGroupSession) error { return nil }
func (h *handler) ConsumeClaim(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		value := string(message.Value)
		fmt.Println(value)
		session.MarkMessage(message, "")
	}
	return nil
}

func main() {
	// Koneksi ke Kafka
	config := sarama.NewConfig()
	config.ClientID = "my-client"
	config.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	config.Consumer.Offsets.Initial = sarama.OffsetOldest

	consumerGroup, err := sarama.NewConsumerGroup([]string{"localhost:9092"}, "my-group", config)
	if err != nil {
		panic(err)
	}
	defer consumerGroup.Close()

	// Subscribe ke topic
	err = consumerGroup.Consume(context.Background(), []string{"my-topic"}, &handler{})
	if err != nil {
		panic(err)
	}
}
