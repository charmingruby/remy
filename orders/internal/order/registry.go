package order

import (
	"github.com/charmingruby/remy-orders/internal/order/contract"
	"github.com/charmingruby/remy-orders/internal/order/service"
	grpcTransport "github.com/charmingruby/remy-orders/internal/order/transport/grpc"
	amqp "github.com/rabbitmq/amqp091-go"
	"google.golang.org/grpc"
)

func NewServiceRegistry(orderRepo contract.OrderRepository) contract.OrderService {
	orderSvc := service.NewServiceRegistry(orderRepo)
	return orderSvc
}

type ServiceRegistry struct{}

func NewGRPCHandler(grpcServer *grpc.Server, orderSvc contract.OrderService, ch *amqp.Channel) {
	grpcTransport.NewGRPCHandler(grpcServer, orderSvc, ch)
}
