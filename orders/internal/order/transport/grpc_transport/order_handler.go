package grpc_transport

import (
	"context"
	"log"

	pb "github.com/charmingruby/remy-common/api"
	"github.com/charmingruby/remy-orders/internal/order/contract"
)

type gRPCOrderHandler struct {
	pb.UnimplementedOrderServiceServer

	service contract.OrderService
}

func (h *gRPCOrderHandler) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Printf("New order received! Order %v\n", p)

	order := &pb.Order{
		ID: "2",
	}

	h.service.CreateOrderService(ctx, p)

	return order, nil
}
