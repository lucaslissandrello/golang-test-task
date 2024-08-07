package consumers

import (
	"context"
	"encoding/json"
	"github.com/go-redis/redis/v8"
	"github.com/streadway/amqp"
	"log"
	"twitch_chat_analysis/cmd/models"
)

type MessageConsumer struct {
	RedisClient *redis.Client
}

func (m *MessageConsumer) ProcessMessages(delivery amqp.Delivery) {
	messageBody := models.MessageBody{}
	err := json.Unmarshal(delivery.Body, &messageBody)
	if err != nil {
		log.Printf("Failed to unmarshal message: %s", err)
		nackErr := delivery.Nack(false, true)
		if nackErr != nil {
			log.Printf("error nacking the pubsub message: %v\n", nackErr)
		}
		return
	}

	messageBodyJSON, err := json.Marshal(messageBody)
	if err != nil {
		log.Printf("Failed to marshal message to JSON: %v", err)
		nackErr := delivery.Nack(false, true)
		if nackErr != nil {
			log.Printf("error nacking the pubsub message: %v\n", nackErr)
		}
		return
	}

	redisSetErr := m.RedisClient.Set(context.Background(), messageBody.Sender, messageBodyJSON, 0).Err()
	if redisSetErr != nil {
		log.Printf("error setting the message to redis db: %v", err)
		nackErr := delivery.Nack(false, true)
		if nackErr != nil {
			log.Printf("error nacking the pubsub message: %v\n", nackErr)
		}
		return
	}

	ackErr := delivery.Ack(false)
	if ackErr != nil {
		log.Printf("error acking the pubsub message: %v\n", ackErr)
	}
	return
}
