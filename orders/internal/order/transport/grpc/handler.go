package grpc

import (
	pb "github.com/charmingruby/remy-common/api"
	"github.com/charmingruby/remy-orders/internal/order/contract"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

func NewGRPCHandler(
	grpcServer *grpc.Server,
	orderSvc contract.OrderService,
	channel *amqp.Channel) {
	orderHandler := &gRPCOrderHandler{
		service: orderSvc,
		ch:      channel,
	}
	pb.RegisterOrderServiceServer(grpcServer, orderHandler)
}
