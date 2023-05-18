package api

import (
	"fmt"
	"net/http"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
)

type CompanyInfoResp struct {
	Id          uuid.UUID `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description,omitempty"`
	Employees   int       `json:"employees"`
	Registered  bool      `json:"registered"`
	Type        string    `json:"type"`
}

func CompanyInfo(req *Req, resp *Resp) {
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

	// Retrieve company info.
	companyTup, err := c.CompanyDb.GetCompanyById(companyId)
	if err != nil {
		fmt.Printf("failed to get company info by Id=%v: %v\n", companyId, err)
		resp.Send(http.StatusNotFound)
		return
	}

	// Construct response.
	respBody := &CompanyInfoResp{
		Id:          companyTup.Id,
		Name:        companyTup.Name,
		Description: companyTup.Description,
		Employees:   companyTup.Employees,
		Registered:  companyTup.Registered,
		Type:        companyTup.Type,
	}

	go func() {
		// Send to Kafka.
		producer, err := kafka.NewProducer(&kafka.ConfigMap{
			"bootstrap.servers": c.kafkaServerUrl,
		})
		if err != nil {
			fmt.Printf("Failed to create producer: %v\n", err)
			return
		}
		defer producer.Close()

		topic := "company_info"
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
	resp.SendData(RC_COMPANY_INFO, respBody)
}
