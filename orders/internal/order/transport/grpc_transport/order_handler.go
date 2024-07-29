package grpc_transport

import (
	"context"
	"log"

	pb "github.com/charmingruby/remy-common/api"
)

type gRPCOrderHandler struct {
	pb.UnimplementedOrderServiceServer
}

func (h *gRPCOrderHandler) CreateOrder(context.Context, *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Println("New order received")

	o := &pb.Order{
		ID: "42",
	}

	return o, nil
}
