package service

import (
	"context"

	pb "github.com/charmingruby/remy-common/api"
	"github.com/google/uuid"
)

func (s *ServiceRegistry) CreateOrderService(ctx context.Context, input *pb.CreateOrderRequest) (*pb.Order, error) {
	items, err := s.ValidateOrderService(ctx, input)
	if err != nil {
		return nil, err
	}

	order := &pb.Order{
		ID:         uuid.NewString(),
		CustomerID: input.CustomerID,
		Items:      items,
		Status:     "pending",
	}

	return order, err
}
