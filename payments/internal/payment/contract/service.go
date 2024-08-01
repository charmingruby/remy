package contract

import (
	"context"

	pb "github.com/charmingruby/remy-common/api"
)

type PaymentService interface {
	CreatePayment(context.Context, *pb.Order) (string, error)
}
