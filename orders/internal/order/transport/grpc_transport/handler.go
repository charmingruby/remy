package grpc_transport

import (
	pb "github.com/charmingruby/remy-common/api"
	"github.com/charmingruby/remy-orders/internal/order/contract"
	"google.golang.org/grpc"
)

func NewGRPCHandler(
	grpcServer *grpc.Server,
	orderSvc contract.OrderService) {
	orderHandler := &gRPCOrderHandler{
		service: orderSvc,
	}
	pb.RegisterOrderServiceServer(grpcServer, orderHandler)
}
