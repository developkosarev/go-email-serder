package main

import (
	"fmt"
	"os"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Printf("Hello v1 \n")
	fmt.Println("Hello v2")

	// Define RabbitMQ server URL.
	amqpServerURL := os.Getenv("AMQP_SERVER_URL")
	amqpServerURL = "amqp://magento:magento@192.168.0.74:5672"
	fmt.Println(amqpServerURL)

	// Create a new RabbitMQ connection.
	connectRabbitMQ, err := amqp.Dial(amqpServerURL)

	if err != nil {
		panic(err)
	}
	defer connectRabbitMQ.Close()

	// Let's start by opening a channel to our RabbitMQ
	// instance over the connection we have already
	// established.
	channelRabbitMQ, err := connectRabbitMQ.Channel()
	if err != nil {
		panic(err)
	}
	defer channelRabbitMQ.Close()

	// With the instance and declare Queues that we can
	// publish and subscribe to.
	_, err = channelRabbitMQ.QueueDeclare(
		"QueueService1", // queue name
		true,            // durable
		false,           // auto delete
		false,           // exclusive
		false,           // no wait
		nil,             // arguments
	)
	if err != nil {
		panic(err)
	}

	// Create a message to publish.
	message := amqp.Publishing{
		ContentType: "text/plain",
		Body:        []byte("Hello"),
	}

	// Attempt to publish a message to the queue.
	err1 := channelRabbitMQ.Publish(
		"",              // exchange
		"QueueService1", // queue name
		false,           // mandatory
		false,           // immediate
		message,         // message to publish
	)

	if err1 != nil {
		panic(err1)
	}
}
