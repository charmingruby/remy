package service

import (
	"context"

	pb "github.com/charmingruby/remy-common/api"
)

func (s *PaymentService) CreatePayment(ctx context.Context, req *pb.Order) (string, error) {
	return "", nil
}
