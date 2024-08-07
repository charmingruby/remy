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
	h.mux.Handle("/", http.FileServer(http.Dir("./static")))
	h.mux.HandleFunc("POST /api/customers/{customerID}/orders", h.handleCreateOrder)
	h.mux.HandleFunc("GET /api/customers/{customerID}/orders/{orderID}", h.handleGetOrder)
}
