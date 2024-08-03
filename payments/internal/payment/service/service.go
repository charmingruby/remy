package service

import "github.com/charmingruby/remy-payments/internal/payment/processor"

func NewPaymentService(processor processor.PaymentProcessor) *PaymentService {
	return &PaymentService{
		Processor: processor,
	}
}

type PaymentService struct {
	Processor processor.PaymentProcessor
}
