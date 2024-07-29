package order

import (
	"github.com/charmingruby/remy-orders/internal/order/contract"
	"github.com/charmingruby/remy-orders/internal/order/service"
	"github.com/charmingruby/remy-orders/internal/order/transport/grpc_transport"
	"google.golang.org/grpc"
)

func NewServiceRegistry(orderRepo contract.OrderRepository) *ServiceRegistry {
	orderSvc := service.NewServiceRegistry(orderRepo)

	return &ServiceRegistry{
		OrderService: orderSvc,
	}
}

type ServiceRegistry struct {
	OrderService contract.OrderService
}

func NewGRPCHandler(grpcServer *grpc.Server, orderSvc contract.OrderService) {
	grpc_transport.NewGRPCHandler(grpcServer, orderSvc)
}
