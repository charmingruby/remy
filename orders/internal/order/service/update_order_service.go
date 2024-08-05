package service

import (
	"context"

	pb "github.com/charmingruby/remy-common/api"
)

func (s *ServiceRegistry) UpdateOrderService(ctx context.Context, input *pb.Order) (*pb.Order, error) {
	err := s.orderRepository.Update(ctx, input.ID, input)
	if err != nil {
		return nil, err
	}

	return input, nil
}
