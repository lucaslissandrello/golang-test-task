package controller

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"net/http"
	"twitch_chat_analysis/cmd/models"
)

type PusherController struct {
	RabbitMQChannel *amqp.Channel
	RabbitKey       string
}

func (p *PusherController) PushMessageToQueue(c *gin.Context) {
	var messageBody models.MessageBody
	if err := c.BindJSON(&messageBody); err != nil {
		c.JSON(http.StatusBadRequest, fmt.Errorf("json body is incorrect: %v", err))
		return
	}

	messageJSON, err := json.Marshal(messageBody)
	if err != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf("error marshalling the message to send to the queue: %v", err))
		return
	}

	rabbitMessage := amqp.Publishing{
		ContentType: "application/json",
		Body:        messageJSON,
	}

	publishErr := p.RabbitMQChannel.Publish("", p.RabbitKey, false, false, rabbitMessage)
	if publishErr != nil {
		c.JSON(http.StatusInternalServerError, fmt.Errorf("error publishing message to the queue the: %v", err))
		return
	}

	c.JSON(http.StatusOK, models.MessageSuccess{Data: "message published successfully"})

}
