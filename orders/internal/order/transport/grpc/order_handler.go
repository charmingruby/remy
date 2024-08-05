package grpc

import (
	"context"
	"encoding/json"
	"log"

	pb "github.com/charmingruby/remy-common/api"
	"github.com/charmingruby/remy-common/broker"
	"github.com/charmingruby/remy-orders/internal/order/contract"
	amqp "github.com/rabbitmq/amqp091-go"
)

type gRPCOrderHandler struct {
	pb.UnimplementedOrderServiceServer

	service contract.OrderService
	ch      *amqp.Channel
}

func (h *gRPCOrderHandler) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Printf("New order received! Order %v\n", p)

	items, err := h.service.ValidateOrderService(ctx, p)
	if err != nil {
		return nil, err
	}

	order, err := h.service.CreateOrderService(ctx, p, items)
	if err != nil {
		return nil, err
	}

	jsonOrder, err := json.Marshal(order)
	if err != nil {
		return nil, err
	}

	q, err := h.ch.QueueDeclare(broker.OrderCreatedEvent, true, false, false, false, nil)
	if err != nil {
		return nil, err
	}

	h.ch.PublishWithContext(ctx, "", q.Name, false, false, amqp.Publishing{
		ContentType:  "application/json",
		Body:         jsonOrder,
		DeliveryMode: amqp.Persistent,
	})

	return order, nil
}

func (h *gRPCOrderHandler) GetOrder(ctx context.Context, req *pb.GetOrderRequest) (*pb.Order, error) {
	return h.service.GetOrderService(ctx, req)
}
