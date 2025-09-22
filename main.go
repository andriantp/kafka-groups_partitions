package main

import (
	"context"
	"flag"
	"fmt"
	"kafka-go/kafka"
	"log"
	"os"

	k "github.com/segmentio/kafka-go"
)

var (
	broker = "kafka:9092"   // localhost:29092
	topic  = "kafka-replay" // adjust by topic
)

func main() {
	flag.Usage = func() {
		log.Printf("Usage:")
		log.Println(" ================== consumer ================== ")
		log.Printf("go run . consumer")
		log.Printf("go run . consumer-group groupName")
		log.Printf("go run . consumer-replay groupName")
		log.Printf("go run . consumer-auto groupName")
		log.Println(" ================== producer ================== ")
		log.Printf("go run . producer key value")
		log.Printf("go run . producer-group")
		flag.PrintDefaults()
	}
	flag.Parse()
	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(1)
	}

	switch flag.Args()[0] {
	case "consumer":
		kafka.Consumer(broker, topic)

	case "consumer-group":
		groupID := flag.Args()[1]
		kafka.ConsumerGroup(broker, "kafka-2", groupID)

	case "producer":
		msg := k.Message{
			Key:   []byte(flag.Args()[1]),
			Value: []byte(flag.Args()[2]),
		}

		if err := kafka.Producer(broker, topic, msg); err != nil {
			log.Fatalf("Producer:%v", err)
		}

	case "producer-group":
		writer := k.NewWriter(k.WriterConfig{
			Brokers:  []string{broker},
			Topic:    topic,
			Balancer: &k.RoundRobin{}, // <-- RoundRobin to all partition
		})
		defer writer.Close()

		fmt.Println("Sending 50 messages...")
		for i := 13; i <= 50; i++ {
			msg := k.Message{
				Key: nil, // for round-robin
				// Key: []byte(fmt.Sprintf("key-%d", i)),
				Value: []byte(fmt.Sprintf("value-%d", i)),
			}

			if err := writer.WriteMessages(context.Background(), msg); err != nil {
				log.Fatalf("Producer error: %v", err)
			}
			fmt.Printf("Sent: value-%d\n", i)
		}
		fmt.Println("All messages sent!")

	case "consumer-replay":
		groupID := flag.Args()[1]
		kafka.ConsumerReplay(broker, topic, groupID)

	case "consumer-auto":
		groupID := flag.Args()[1]
		kafka.ConsumerAutoOffset(broker, topic, groupID)

	}

}
