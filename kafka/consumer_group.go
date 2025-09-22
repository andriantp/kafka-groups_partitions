package kafka

import (
	"context"
	"fmt"
	"log"

	k "github.com/segmentio/kafka-go"
)

// ConsumerGroup membaca pesan dari Kafka dengan GroupID
func ConsumerGroup(broker, topic, groupID string) {
	reader := k.NewReader(k.ReaderConfig{
		Brokers:  []string{broker},
		Topic:    topic,
		GroupID:  groupID, // inilah yang bikin dia masuk ke consumer group
		MinBytes: 1,
		MaxBytes: 10e6,
	})
	defer reader.Close()

	fmt.Printf("Consumer Group [%s] started...\n", groupID)

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("Error reading message:", err)
		}
		log.Printf("[%s] partition=%d offset=%d key=%s value=%s\n",
			groupID, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
	}
}
