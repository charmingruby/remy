package contract

import (
	"context"

	pb "github.com/charmingruby/remy-common/api"
)

type OrderRepository interface {
	Create(context.Context, *pb.CreateOrderRequest, []*pb.Item) (string, error)
	Get(ctx context.Context, id, customerID string) (*pb.Order, error)
	Update(ctx context.Context, id string, input *pb.Order) error
}
