package order

import (
	"github.com/charmingruby/remy-orders/internal/order/contract"
	"github.com/charmingruby/remy-orders/internal/order/service"
	"github.com/charmingruby/remy-orders/internal/order/transport/grpc_transport"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

func NewServiceRegistry(orderRepo contract.OrderRepository) contract.OrderService {
	orderSvc := service.NewServiceRegistry(orderRepo)
	return orderSvc
}

type ServiceRegistry struct{}

func NewGRPCHandler(grpcServer *grpc.Server, orderSvc contract.OrderService, ch *amqp.Channel) {
	grpc_transport.NewGRPCHandler(grpcServer, orderSvc, ch)
}
