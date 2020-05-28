package messaging

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func StartAsync(queue string, builder Builder) error {
	connectionFactory := RabbitConnectionFactory{}
	channel, err := connectionFactory.CreateChannel()
	if err != nil {
		log.Fatal(err)
	}
	defer channel.Close()

	messages, err := channel.Consume(
		queue, // queue
		"",               // consumer
		false,             // auto-ack
		false,            // exclusive
		false,            // no-local
		false,            // no-wait
		nil,              // args
	)
	if err != nil {
		log.Fatal(err)
	}

	forever := make(chan bool)
	go func() {
		for message := range messages {

			HandleMessage(message, builder)
			_ = message.Ack(false)
		}
	}()
	<-forever
	return err
}

func HandleMessage(message amqp.Delivery, builder Builder) {
	messageType := message.Headers["MessageType"]
	if messageType == nil {
		return
	}
	str := fmt.Sprintf("%v", messageType)
	handler := builder.TryGetValue(str)
	if handler == nil {
		return
	}

	go handler.HandleMessageAsync(message.Body)
}
