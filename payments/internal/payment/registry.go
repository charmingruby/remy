package payment

import (
	"net/http"

	"github.com/charmingruby/remy-payments/internal/payment/contract"
	"github.com/charmingruby/remy-payments/internal/payment/processor"
	"github.com/charmingruby/remy-payments/internal/payment/service"
	"github.com/charmingruby/remy-payments/internal/payment/transport/rest"
	amqp "github.com/rabbitmq/amqp091-go"
)

func NewPaymentService(processor processor.PaymentProcessor) contract.PaymentService {
	paymentSvc := service.NewPaymentService(processor)
	return paymentSvc
}

func NewPaymentHTTPHandler(mux *http.ServeMux, ch *amqp.Channel, stripeEndpointSecret string) rest.Handler {
	return *rest.NewHTTPHandler(mux, ch, stripeEndpointSecret)
}
