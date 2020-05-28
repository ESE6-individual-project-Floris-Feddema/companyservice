package messaging

import (
	"github.com/streadway/amqp"
)

type RabbitConnectionFactory struct{
}

var connection *amqp.Connection


func (connectionFactory RabbitConnectionFactory) GetConnection(connectionString string) (*amqp.Connection, error) {
	if connection != nil {
		 return connection, nil
	}

	var err error
	connection, err = amqp.Dial(connectionString)

	if err != nil {
		return nil, err
	}
	return connection, err
}

func (connectionFactory RabbitConnectionFactory) CreateChannel() (*amqp.Channel, error) {
	var ch, err = connection.Channel()
	if err != nil {
		return nil, err
	}
	return ch, nil
}