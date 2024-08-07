package broker

import (
	"fmt"
	"github.com/streadway/amqp"
)

func ConnectBroker(brokerURL string) (*amqp.Connection, error) {
	connection, err := amqp.Dial(fmt.Sprintf("amqp://guest:guest@%s", brokerURL))
	if err != nil {
		fmt.Printf("error opening the connection to the broker")
		return nil, err
	}
	defer connection.Close()

	return connection, nil
}

func CreateConsumer(queueName string, brokerConnection amqp.Connection) (<-chan amqp.Delivery, error) {
	ch, err := brokerConnection.Channel()
	if err != nil {
		return nil, fmt.Errorf("error creating a channel from the broker connection")
	}
	defer ch.Close()

	return ch.Consume(queueName, "", true, false, false, false, nil)
}
