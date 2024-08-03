package payment

import (
	"github.com/charmingruby/remy-payments/internal/payment/contract"
	"github.com/charmingruby/remy-payments/internal/payment/processor"
	"github.com/charmingruby/remy-payments/internal/payment/service"
)

func NewPaymentService(processor processor.PaymentProcessor) contract.PaymentService {
	paymentSvc := service.NewPaymentService(processor)
	return paymentSvc
}
