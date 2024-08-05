package service

import (
	"github.com/charmingruby/remy-payments/internal/gateway"
	"github.com/charmingruby/remy-payments/internal/payment/processor"
)

func NewPaymentService(
	processor processor.PaymentProcessor,
	ordersGw gateway.OrdersGateway) *PaymentService {
	return &PaymentService{
		Processor:     processor,
		OrdersGateway: ordersGw,
	}
}

type PaymentService struct {
	Processor     processor.PaymentProcessor
	OrdersGateway gateway.OrdersGateway
}
