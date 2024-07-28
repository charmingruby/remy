package rest

import "net/http"

func NewHandler(mux *http.ServeMux) *Handler {
	return &Handler{
		mux: mux,
	}
}

type Handler struct {
	mux *http.ServeMux
	// gateway
}

func (h *Handler) Register() {
	h.mux.HandleFunc("POST /api/customers/{customerID}/orders", h.handleCreateOrder)
}
