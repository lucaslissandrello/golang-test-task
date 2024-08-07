package controller

import (
	"github.com/gin-gonic/gin"
	"twitch_chat_analysis/cmd/services"
)

type MessagesController struct {
	MessagesService services.MessageService
}

func (m *MessagesController) GetMessages(c *gin.Context) {
	//TODO parse sender and receiver and call the service
}
