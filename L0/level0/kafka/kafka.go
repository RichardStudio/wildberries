package kafka

import (
	"context"
	"encoding/json"
	"level0/database"
	"level0/models"
	"log"

	"github.com/segmentio/kafka-go"
)

func ConsumeKafkaMessages(dbClient *database.DatabaseClient) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{"localhost:9092"},
		Topic:   "orders",
		GroupID: "order_service_group",
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("message at offset %d: %s = %s\n", m.Offset, string(m.Key), string(m.Value))

		var order models.Order
		if err := json.Unmarshal(m.Value, &order); err != nil {
			log.Printf("Error unmarshalling message: %v", err)
			continue
		}

		database.Cache[order.OrderUID] = order
		dbClient.SaveOrder(order)
	}
}
