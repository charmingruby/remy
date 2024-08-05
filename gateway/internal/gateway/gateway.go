package gateway

import (
	"context"

	pb "github.com/charmingruby/remy-common/api"
)

type OrdersGateway interface {
	CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.Order, error)
	GetOrder(ctx context.Context, orderID, customerID string) (*pb.Order, error)
}
