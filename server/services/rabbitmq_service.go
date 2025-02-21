package services

import (
	"context"
	"fmt"
	"log"
	"sync"
	"time"

	"github.com/overal-x/formatio/config"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/samber/do"
)

type IRabbitMQService interface {
	Publish(PublishArgs) error
	Subscribe(SubscribeArgs) error
	SubscribeWithWorkers(int, SubscribeArgs)
}

type RabbitMQService struct {
	connection *amqp.Connection
}

type PublishArgs struct {
	Queue   string
	Content string
}

type SubscribeArgs struct {
	Queue    string
	Callback func(content string) error
}

func (r *RabbitMQService) Publish(args PublishArgs) error {
	ch, err := r.connection.Channel()
	if err != nil {
		log.Println("[x]", err)

		return err
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		args.Queue, // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		log.Println("[x]", err)

		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = ch.PublishWithContext(ctx,
		"",     // exchange
		q.Name, // routing key
		false,  // mandatory
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "text/plain",
			Body:         []byte(args.Content),
		})
	if err != nil {
		log.Println("[x]", err)

		return err
	}

	return nil
}

func (r *RabbitMQService) Subscribe(args SubscribeArgs) error {
	ch, err := r.connection.Channel()
	if err != nil {
		log.Println("[x]", err)

		return err
	}

	defer ch.Close()

	q, err := ch.QueueDeclare(
		args.Queue, // name
		true,       // durable
		false,      // delete when unused
		false,      // exclusive
		false,      // no-wait
		nil,        // arguments
	)
	if err != nil {
		log.Printf("[x][%s] %s", args.Queue, err)

		return err
	}

	err = ch.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)
	if err != nil {
		return err
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		false,  // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		return err
	}

	var forever chan struct{}

	go func() {
		for d := range msgs {
			err := args.Callback(string(d.Body))
			if err != nil {
				log.Printf("[*][%s] Error %s", args.Queue, err)
				break
			}

			d.Ack(false)
		}
	}()

	log.Printf("[*][%s] Waiting for messages. To exit press CTRL+C", args.Queue)
	<-forever

	return nil
}

func (r *RabbitMQService) SubscribeWithWorkers(workers int, args SubscribeArgs) {
	var wg sync.WaitGroup

	log.Printf("[x][%s] Spawning %d workers", args.Queue, workers)

	for i := 1; i <= workers; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			r.Subscribe(args)
		}()

		log.Printf("[x][%s] Worker %d just joined the party", args.Queue, i)
	}

	wg.Wait()
}

func NewRabbitMQService(i *do.Injector) (IRabbitMQService, error) {
	connection := do.MustInvoke[*amqp.Connection](i)

	return &RabbitMQService{connection: connection}, nil
}

func NewRabbitMQConnection(i *do.Injector) (*amqp.Connection, error) {
	env := do.MustInvoke[*config.Env](i)
	connection, err := amqp.Dial(env.RABBITMQ_URL)

	if err != nil {
		return nil, fmt.Errorf("[x] %s: while connecting", err)
	}

	log.Println("[x] connection established")

	return connection, nil
}
