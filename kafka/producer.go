package kafka

import (
	"context"
	"fmt"

	k "github.com/segmentio/kafka-go"
)

func Producer(broker, topic string, msg k.Message) error {
	writer := k.NewWriter(k.WriterConfig{
		Brokers: []string{broker}, // ganti sesuai port Docker Kafka
		Topic:   topic,
	})
	defer writer.Close()
	fmt.Println("NewWriter successfully!")

	err := writer.WriteMessages(context.Background(), msg)
	if err != nil {
		return err
	}

	fmt.Println("Message sent successfully!")
	return nil
}
