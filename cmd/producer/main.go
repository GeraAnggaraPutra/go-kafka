package main

import (
	"github.com/IBM/sarama"
)

func main() {
	// Koneksi ke Kafka
	config := sarama.NewConfig()
	config.ClientID = "my-client"

	// Buat pesan
	message := &sarama.ProducerMessage{
		Topic: "my-topic",
		Value: sarama.StringEncoder("Hello, world!"),
	}

	// Kirim pesan
	p, err := sarama.NewAsyncProducer([]string{"localhost:9092"}, config)
	if err != nil {
		panic(err)
	}
	defer p.Close()

	p.Input() <- message // Kirim pesan ke channel input
}
