package main

import (
	"fmt"
	"listener/event"
	"log"
	"math"
	"os"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// try to conncect to rabbitmq
	rabbitConn, err := connect()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	defer rabbitConn.Close()
	

	// start listenening for messages
	log.Println("Listeneing for and consuming RabbitMQ messages")

	// create a consumer
	consumer, err := event.NewConsumer(rabbitConn)
	if err != nil {
		panic(err)
	}

	// watch the queue and consume event from topic
	err = consumer.Listen([]string{"log.INFO", "log.WARNING", "log.ERROR"})
	if err != nil {
		log.Println(err)
	}
}

func connect() (*amqp.Connection, error) {
	var counts int64
	var backOff = 1 * time.Second
	var connection *amqp.Connection

	// don't continue uintil rabbit is ready
	for {
		c, err := amqp.Dial("amqp://rabbitmq:password@localhost")
		if err != nil {
			fmt.Println("RabbitMQ not yet ready...")
			counts++
		} else {
			log.Println("Connected to RabbitMQ!")
			connection = c
			break
		}

		// retry not more that 5 times
		if counts > 5 {
			fmt.Println(err)
			return nil, err
		}

		// delay exponentially aftet first try to connect; 1,2,4,8,16
		backOff = time.Duration(math.Pow(float64(counts), 2)) * time.Second
		log.Println("backing off")
		time.Sleep(backOff)
		continue
	}

	return connection, nil
}
