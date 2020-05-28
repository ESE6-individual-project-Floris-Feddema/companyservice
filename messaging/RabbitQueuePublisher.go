package messaging

import (
	"encoding/json"
	"github.com/streadway/amqp"
	"log"
)

func(publisher MessagePublisher) PublishMessageAsync(exchange string, routingKey string, messageType string, message interface{}) error {
	connectionFactory := RabbitConnectionFactory{}
	channel, err := connectionFactory.CreateChannel()
	if err != nil {
		log.Fatal(err)
	}

	headers := amqp.Table{}
	headers["MessageType"] = messageType

	body, err := json.Marshal(message)
	if err != nil {
		log.Fatal(err)
	}

	err = channel.Publish(exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: 2,
			Headers:      headers,
			ContentType:  "application/json",
			Body:         body,
		})
	if err != nil {
		log.Fatal(err)
	}

	return nil
}