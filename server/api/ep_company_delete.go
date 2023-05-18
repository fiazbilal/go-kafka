package api

import (
	"fmt"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
)

func CompanyDelete(req *Req, resp *Resp) {
	// get query params.
	qVals := req.URL.Query()

	// Parse out id variable.
	companyIdStr := qVals.Get("uuid")
	companyId, err := uuid.Parse(companyIdStr)
	if err != nil {
		fmt.Printf(
			"failed to parse companyIdStr=%s: %v\n",
			companyIdStr, err,
		)
		resp.Send(http.StatusBadRequest)
		return
	}

	// Company create.
	err = c.CompanyDb.DeleteCompany(companyId)
	if err != nil {
		fmt.Printf("failed to delete company by Id=%v: %v\n", companyId, err)
		resp.Send(http.StatusInternalServerError)
		return
	}

	go func() {
		// Send to Kafka.
		producer, err := kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers": "localhost:9092",
		})
		if err != nil {
			fmt.Printf("Failed to create producer: %v\n", err)
			return
		}
		defer producer.Close()

		topic := "company_deleted"
		msg := &kafka.Message{
			TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
			Value:          []byte(companyId.String()),
		}
		err = producer.Produce(msg, nil)
		if err != nil {
			fmt.Printf("Failed to produce message: %v\n", err)
			return
		}
	}()

	// Send.
	resp.Send(RC_COMPANY_DELETE)
}
