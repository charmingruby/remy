package order

import (
	"github.com/charmingruby/remy-orders/internal/order/domain"
	"github.com/charmingruby/remy-orders/internal/order/transport/grpc_transport"
	"google.golang.org/grpc"
)

func NewServiceRegistry(svc domain.Service) *ServiceRegistry {
	return &ServiceRegistry{
		service: svc,
	}
}

type ServiceRegistry struct {
	service domain.Service
}

func NewGRPCHandler(grpcServer *grpc.Server) {
	grpc_transport.NewGRPCHandler(grpcServer)
}
