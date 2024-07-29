package gateway

import (
	"context"

	pb "github.com/charmingruby/remy-common/api"
)

type OrdersGateway interface {
	CreateOrderService(ctx context.Context, req *pb.CreateOrderRequest) (*pb.Order, error)
}
