package processor

import pb "github.com/charmingruby/remy-common/api"

type PaymentProcessor interface {
	CreatePaymentLink(*pb.Order) (string, error)
}
