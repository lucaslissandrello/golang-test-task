package models

type MessageSuccess struct {
	Data string `json:"data"`
}

type MessageBody struct {
	Sender   string `json:"sender"`
	Receiver string `json:"receiver"`
	Message  string `json:"message"`
}
