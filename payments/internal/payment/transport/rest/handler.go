package rest

import (
	"net/http"

	amqp "github.com/rabbitmq/amqp091-go"
)

func NewHTTPHandler(mux *http.ServeMux, ch *amqp.Channel, stripeEndpointSecret string) *Handler {
	return &Handler{
		mux:                  mux,
		ch:                   ch,
		stripeEndpointSecret: stripeEndpointSecret,
	}
}

type Handler struct {
	mux                  *http.ServeMux
	ch                   *amqp.Channel
	stripeEndpointSecret string
}

func (h *Handler) Register() {
	h.mux.HandleFunc("POST /webhook", h.checkoutWebhookEndpoint)
}
