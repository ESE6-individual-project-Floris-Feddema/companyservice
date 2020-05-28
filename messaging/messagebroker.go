package messaging

import (
	"companyservice/utils"
	"github.com/streadway/amqp"
	"log"
)

func InitMessageBroker() error {

	conn, err := amqp.Dial(utils.EnvVar("RABBITMQ_CONNECTION_STRING"))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatal(err)
	}
	defer ch.Close()

	msgs, err := ch.Consume(
		"companyservice", // queue
		"",               // consumer
		true,             // auto-ack
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
		for d := range msgs {
			log.Printf(" [x] %s", d.Body)
		}
	}()
	<-forever

	return err
}
