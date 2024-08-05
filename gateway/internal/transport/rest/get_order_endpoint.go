package rest

import (
	"net/http"

	common "github.com/charmingruby/remy-common"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) handleGetOrder(w http.ResponseWriter, r *http.Request) {
	customerID := r.PathValue("customerID")
	orderID := r.PathValue("orderID")

	o, err := h.ordersGateway.GetOrder(r.Context(), orderID, customerID)
	rStatus := status.Convert(err)
	if rStatus != nil {
		if rStatus.Code() != codes.InvalidArgument {
			common.WriteError(w, http.StatusBadRequest, rStatus.Message())
			return
		}

		common.WriteError(w, http.StatusInternalServerError, rStatus.Message())
		return
	}

	common.WriteJSON(w, http.StatusOK, o)
}
