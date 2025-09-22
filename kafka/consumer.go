package kafka

import (
	"context"
	"fmt"
	"log"

	k "github.com/segmentio/kafka-go"
)

func Consumer(broker, topic string) {
	reader := k.NewReader(k.ReaderConfig{
		Brokers:   []string{broker}, // ganti sesuai port Docker Kafka
		Topic:     topic,
		Partition: 0,
		MinBytes:  1,
		MaxBytes:  10e6,
	})
	defer reader.Close()

	fmt.Println("Starting to read messages...")
	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("Error reading message:", err)
		}
		log.Printf("Received: key=%s value=%s\n", string(msg.Key), string(msg.Value))
	}
}
