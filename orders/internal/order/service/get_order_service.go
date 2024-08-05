package service

import (
	"context"
	"fmt"

	pb "github.com/charmingruby/remy-common/api"
)

func (s *ServiceRegistry) GetOrderService(ctx context.Context, input *pb.GetOrderRequest) (*pb.Order, error) {
	order, err := s.orderRepository.Get(ctx, input.OrderID, input.CustomerID)
	if err != nil {
		return nil, fmt.Errorf("order not found: %s", input.OrderID)
	}

	return order, nil
}
