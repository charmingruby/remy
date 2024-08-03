package payment

import (
	"github.com/charmingruby/remy-payments/internal/payment/contract"
	"github.com/charmingruby/remy-payments/internal/payment/service"
)

func NewPaymentService() contract.PaymentService {
	paymentSvc := service.NewPaymentService()
	return paymentSvc
}
