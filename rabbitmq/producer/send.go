package main

import (
	"github.com/streadway/amqp"
	"log"
	"os"
	"strings"
)

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}

func bodyFrom(args []string) string {
	var s string
	if (len(args) < 2) || os.Args[1] == "" {
		s = "hello"
	} else {
		s = strings.Join(args[1:], " ")
	}
	return s
}

func main() {

	//conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	//failOnError(err, "Failed to connect to RabbitMQ")
	//defer conn.Close()
	//
	//ch, err := conn.Channel()
	//failOnError(err, "Failed to open a channel")
	//defer ch.Close()
	//
	//q, err := ch.QueueDeclare(
	//	"work", // name
	//	true,   // durable
	//	false,   // delete when unused
	//	false,   // exclusive
	//	false,   // no-wait
	//	nil,     // arguments
	//)
	//failOnError(err, "Failed to declare a queue")
	//
	//body := bodyFrom(os.Args)
	//
	//for i:=0; i<100; i++ {
	//	err = ch.Publish(
	//		"",           // exchange
	//		q.Name,       // routing key
	//		false,        // mandatory
	//		false,
	//		amqp.Publishing {
	//			DeliveryMode: amqp.Persistent,
	//			ContentType:  "text/plain",
	//			Body:         []byte(body),
	//		})
	//	failOnError(err, "Failed to publish a message")
	//	log.Printf(" [x] Sent %s", body)
	//}

	conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
	failOnError(err, "Failed to connect to RabbitMQ")
	defer conn.Close()

	ch, err := conn.Channel()
	failOnError(err, "Failed to open a channel")
	defer ch.Close()

	q, err := ch.QueueDeclare(
		"hello", // name
		true,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")

	body := "Hello World!"
	err = ch.Publish(
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
			ContentType: "text/plain",
			Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message")
}





