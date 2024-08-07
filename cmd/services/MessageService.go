package services

import (
	"context"
	"github.com/go-redis/redis/v8"
	"twitch_chat_analysis/cmd/models"
)

type MessageService struct {
	RedisClient *redis.Client
}

func (m *MessageService) GetMessage(sender, receiver string) ([]models.MessageBody, error) {
	var messagesSender []models.MessageBody
	getErr := m.RedisClient.GetSet(context.Background(), sender, messagesSender).Err()
	if getErr != nil {
		return nil, getErr
	}

	var messagesReceiver []models.MessageBody
	getErr = m.RedisClient.GetSet(context.Background(), receiver, messagesReceiver).Err()
	if getErr != nil {
		return nil, getErr
	}

	//TODO do an inner join between both arrays and order results in chronological order

	return nil, nil
}
