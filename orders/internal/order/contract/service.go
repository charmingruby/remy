package contract

import (
	"context"

	pb "github.com/charmingruby/remy-common/api"
)

type OrderService interface {
	CreateOrderService(ctx context.Context, input *pb.CreateOrderRequest) (*pb.Order, error)
	ValidateOrderService(ctx context.Context, input *pb.CreateOrderRequest) ([]*pb.Item, error)
}
