package rest

import (
	"net/http"

	"github.com/charmingruby/remy-gateway/internal/gateway"
)

func NewHandler(mux *http.ServeMux, ordersGateway gateway.OrdersGateway) *Handler {
	return &Handler{
		mux:           mux,
		ordersGateway: ordersGateway,
	}
}

type Handler struct {
	mux           *http.ServeMux
	ordersGateway gateway.OrdersGateway
}

func (h *Handler) Register() {
	h.mux.HandleFunc("POST /api/customers/{customerID}/orders", h.handleCreateOrder)
}
