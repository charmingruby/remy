package service

import (
	"context"

	pb "github.com/charmingruby/remy-common/api"
)

func (s *PaymentService) CreatePayment(ctx context.Context, req *pb.Order) (string, error) {
	link, err := s.Processor.CreatePaymentLink(req)
	if err != nil {
		return "", err
	}

	// update the order with the link

	return link, nil
}
