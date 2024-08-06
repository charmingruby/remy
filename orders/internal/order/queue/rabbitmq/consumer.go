package rabbitmq

import (
	"context"
	"encoding/json"
	"log"

	pb "github.com/charmingruby/remy-common/api"
	"github.com/charmingruby/remy-common/broker"
	"github.com/charmingruby/remy-orders/internal/order/contract"
	amqp "github.com/rabbitmq/amqp091-go"
)

func NewConsumer(orderSvc contract.OrderService) *Consumer {
	return &Consumer{
		OrderService: orderSvc,
	}
}

type Consumer struct {
	OrderService contract.OrderService
}

func (c *Consumer) Listen(ch *amqp.Channel) {
	q, err := ch.QueueDeclare("", true, false, true, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	err = ch.QueueBind(q.Name, "", broker.OrderPaidEvent, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	msgs, err := ch.Consume(q.Name, "", false, false, false, false, nil)
	if err != nil {
		log.Fatal(err)
	}

	forever := make(chan struct{})

	go func() {
		for d := range msgs {
			log.Printf("Received message: %s", string(d.Body))

			order := &pb.Order{}
			if err := json.Unmarshal(d.Body, order); err != nil {
				d.Nack(false, false)
				log.Printf("failed to unmarshal order: %v", d)
				continue
			}

			_, err := c.OrderService.UpdateOrderService(context.Background(), order)
			if err != nil {
				log.Printf("failed to update order: %v", order)

				if err := broker.HandleRetry(ch, &d); err != nil {
					log.Printf("Error handling retry: %v", err)
				}

				d.Nack(false, false)

				continue
			}

			log.Printf("Order have been updated")
			d.Ack(false)
		}
	}()

	<-forever
}
