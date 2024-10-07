package producer

import (
	util "company/libgo"
	"company/server"
	"fmt"
	"log"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

type Req struct {
	server.Req
}

type Resp struct {
	server.Resp
}

func Main() {
	serverUrl := util.MustOsGetEnv("KAFKA_SERVER_URL")

	log.Printf("Kafka producer connected to: %s", serverUrl)

	// Set up Kafka producer configuration
	producer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": serverUrl,
	})
	if err != nil {
		fmt.Printf("Failed to create producer: %v\n", err)
		return
	}
	defer producer.Close()

	// Create a delivery channel to receive the report of message delivery status
	deliveryChan := make(chan kafka.Event)
	defer close(deliveryChan)

	// Infinite loop to produce messages every 2 seconds
	for i := 1; ; i++ {
		// Create the message payload
		msg := fmt.Sprintf(`{"cloud_id": %d, "action": "cloud_updated"}`, i)
		event := "cloud_events"
		// Produce the message
		err = producer.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &event, Partition: kafka.PartitionAny},
			Value:          []byte(msg),
		}, deliveryChan)

		if err != nil {
			fmt.Printf("Failed to produce message: %v\n", err)
			continue
		}

		// Wait for message delivery report
		e := <-deliveryChan
		m := e.(*kafka.Message)

		if m.TopicPartition.Error != nil {
			fmt.Printf("Failed to deliver message: %v\n", m.TopicPartition.Error)
		} else {
			fmt.Printf("Message delivered to topic %s [partition %d] at offset %v\n",
				*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
		}

		// Sleep for 2 seconds before sending the next message
		time.Sleep(2 * time.Second)
	}

}
