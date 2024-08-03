package rabbitmq

import (
	"context"
	"encoding/json"
	"log"

	pb "github.com/charmingruby/remy-common/api"
	"github.com/charmingruby/remy-common/broker"
	"github.com/charmingruby/remy-payments/internal/payment/contract"
	amqp "github.com/rabbitmq/amqp091-go"
)

func NewConsumer(paymentSvc contract.PaymentService) *Consumer {
	return &Consumer{
		PaymentService: paymentSvc,
	}
}

type Consumer struct {
	PaymentService contract.PaymentService
}

func (c *Consumer) Listen(ch *amqp.Channel) {
	q, err := ch.QueueDeclare(broker.OrderCreatedEvent, true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	forever := make(chan struct{})

	go func() {
		for d := range msgs {
			log.Printf("Received Order: %s", string(d.Body))

			order := &pb.Order{}
			if err := json.Unmarshal(d.Body, order); err != nil {
				log.Printf("failed to unmarshal order: %v", d)
				continue
			}

			paymentLink, err := c.PaymentService.CreatePayment(context.Background(), order)
			if err != nil {
				log.Printf("failed to create payment: %v", err)
				continue
			}

			log.Printf("Payment link created %s", paymentLink)
		}
	}()

	<-forever
}
