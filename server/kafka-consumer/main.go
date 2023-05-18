package consumer

import (
	util "company/libgo"
	"company/server"
	"fmt"
	"log"

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
	groupId := util.MustOsGetEnv("GROUP_ID")

	log.Printf("kafka consumer listening on: %s", serverUrl)

	// Set up Kafka consumer configuration
	consumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": serverUrl,
		"group.id":          groupId,
		"auto.offset.reset": "earliest",
	})
	// Close the consumer
	defer consumer.Close()

	if err != nil {
		fmt.Printf("Failed to create consumer: %v\n", err)
		return
	}

	// Subscribe to topics
	consumer.SubscribeTopics([]string{"company_created", "company_deleted", "company_updated", "company_info"}, nil)

	// Start consuming messages
	for {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			fmt.Printf("Received message from topic '%s': %s\n", *msg.TopicPartition.Topic, string(msg.Value))
		} else {
			fmt.Printf("Error: %v\n", err)
		}
	}

}
