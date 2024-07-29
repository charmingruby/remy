package rest_transport

import (
	"net/http"

	"github.com/charmingruby/remy-gateway/internal/transport/grpc_transport"
)

func NewHandler(mux *http.ServeMux, grpcHandler *grpc_transport.Handler) *Handler {
	return &Handler{
		mux:         mux,
		grpcHandler: grpcHandler,
	}
}

type Handler struct {
	mux         *http.ServeMux
	grpcHandler *grpc_transport.Handler
}

func (h *Handler) Register() {
	h.mux.HandleFunc("POST /api/customers/{customerID}/orders", h.handleCreateOrder)
}
