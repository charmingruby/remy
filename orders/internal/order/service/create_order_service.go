package service

import (
	"context"

	pb "github.com/charmingruby/remy-common/api"
)

func (s *ServiceRegistry) CreateOrderService(
	ctx context.Context,
	input *pb.CreateOrderRequest,
	items []*pb.Item,
) (*pb.Order, error) {
	id, err := s.orderRepository.Create(ctx, input, items)

	order := &pb.Order{
		ID:         id,
		CustomerID: input.CustomerID,
		Status:     "pending",
		Items:      items,
	}

	return order, err
}
