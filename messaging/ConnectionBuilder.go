package messaging

import (
	"log"
)

func AddMessageConsumer(connectionUrl string, queueName string, handlers map[string]MessageHandler) error {
	messageBuilder := Builder{
		messageHandlers: map[string]MessageHandler{},
	}
	for key, element := range handlers {
		log.Printf("[AMQP] Found handler for %v", key)
		messageBuilder.WithHandler(element, key)
	}

	factory := RabbitConnectionFactory{}
	_, err := factory.GetConnection(connectionUrl)
	if err != nil {
		log.Fatal(err)
	}

	go StartAsync(queueName, messageBuilder)
	return err
}

var messagePublisher MessagePublisher

func AddMessagePublisher(connectionUrl string) error {
	messagePublisher = MessagePublisher{}
	factory := RabbitConnectionFactory{}
	_, err := factory.GetConnection(connectionUrl)
	if err != nil {
		log.Fatal(err)
	}
	return nil
}

type MessagePublisher struct {}
type MessageConsumer struct {}


