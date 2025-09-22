package kafka

import (
	"context"
	"fmt"
	"log"
	"time"

	k "github.com/segmentio/kafka-go"
)

// ConsumerReplay membaca pesan dari Kafka dengan manual offset commit
func ConsumerReplay(broker, topic string, groupID string) {
	reader := k.NewReader(k.ReaderConfig{
		Brokers:        []string{broker},
		Topic:          topic,
		GroupID:        groupID,
		MinBytes:       1,
		MaxBytes:       10e6,
		CommitInterval: 0,             // 0 = manual commit
		StartOffset:    k.FirstOffset, // mulai dari offset awal
	})
	defer reader.Close()

	fmt.Printf("Consumer Replay [%s] started...\n", groupID)

	for {
		msg, err := reader.FetchMessage(context.Background())
		if err != nil {
			log.Fatalf("Error fetching message: %v", err)
		}

		// Proses pesan
		log.Printf("[%s] partition=%d offset=%d key=%s value=%s\n",
			groupID, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))

		// Commit offset secara manual setelah selesai memproses message
		if err := reader.CommitMessages(context.Background(), msg); err != nil {
			log.Fatalf("Error committing offset: %v", err)
		}

		// Optional: jeda agar output lebih mudah dibaca
		time.Sleep(100 * time.Millisecond)
	}
}
