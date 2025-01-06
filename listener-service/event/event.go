package event

import (
	amqp "github.com/rabbitmq/amqp091-go"
)

func declareExchange(ch *amqp.Channel) error {
	return ch.ExchangeDeclare(
		"logs_topics", // name of the excahnge
		"topic",      // type
		true,         // durable?
		false,        // auto-deleted?
		false,        // internal useage ?
		false,        // no wait?
		nil,          // arguments?
	)
}

func declareRandomQueue(ch *amqp.Channel) (amqp.Queue, error) {
	return ch.QueueDeclare(
		"",    // name
		false, // durable
		false, // delete when unused?
		true,  // esclusive?
		false, // no wait?
		nil,   // arguments
	)
}
