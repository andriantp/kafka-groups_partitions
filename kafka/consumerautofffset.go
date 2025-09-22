package kafka

import (
	"context"
	"fmt"
	"log"

	k "github.com/segmentio/kafka-go"
)

// ConsumerAutoOffset membaca pesan dari Kafka dengan commit offset otomatis
func ConsumerAutoOffset(broker, topic, groupID string) {
	reader := k.NewReader(k.ReaderConfig{
		Brokers:        []string{broker},
		Topic:          topic,
		GroupID:        groupID,
		MinBytes:       1,
		MaxBytes:       10e6,
		CommitInterval: 1 * 1000, // 1 detik, commit otomatis setiap interval
	})
	defer reader.Close()

	fmt.Printf("Consumer Auto Offset [%s] started...\n", groupID)

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("Error reading message: %v", err)
		}

		// Proses pesan
		log.Printf("[%s] partition=%d offset=%d key=%s value=%s\n",
			groupID, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
	}
}
