package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"twitch_chat_analysis/cmd/broker"
	"twitch_chat_analysis/cmd/consumers"
	controller "twitch_chat_analysis/cmd/controllers"
	"twitch_chat_analysis/cmd/redis"
	"twitch_chat_analysis/cmd/services"
)

func main() {
	r := gin.Default()

	redisClient := redis.ConnectRedis("localhost", "", "user", "password", 0)
	pushController := controller.PusherController{
		RabbitMQChannel: nil, //TODO create rabbitMQChannel
		RabbitKey:       "RabbitKey",
	}

	messageController := controller.MessagesController{MessagesService: services.MessageService{RedisClient: redisClient}}

	MessageConsumer := consumers.MessageConsumer{
		RedisClient: redisClient,
	}

	queueName := "queueName"

	brokerConnection, connectErr := broker.ConnectBroker("localhost:5672")
	if connectErr != nil {
		log.Fatal("error connecting to broker")
	}

	consumers, consumerErr := broker.CreateConsumer(queueName, *brokerConnection)
	if consumerErr != nil {
		log.Fatalf("error creating consumer to queue: %s", queueName)
	}

	go func() { //TODO See how to stop this
		for message := range consumers {
			MessageConsumer.ProcessMessages(message)
		}
	}()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, "worked")
	})

	r.POST("/message", pushController.PushMessageToQueue)
	r.GET("/message/list", messageController.GetMessages)
	r.Run()
}
