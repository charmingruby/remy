package contract

import (
	"context"

	pb "github.com/charmingruby/remy-common/api"
)

type OrderService interface {
	CreateOrderService(ctx context.Context, input *pb.CreateOrderRequest) error
	ValidateOrderService(ctx context.Context, input *pb.CreateOrderRequest) error
}
