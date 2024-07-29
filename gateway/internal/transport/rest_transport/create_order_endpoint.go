package rest_transport

import (
	"net/http"

	common "github.com/charmingruby/remy-common"
	pb "github.com/charmingruby/remy-common/api"
)

func (h *Handler) handleCreateOrder(w http.ResponseWriter, r *http.Request) {
	customerID := r.PathValue("customerID")

	var items []*pb.ItemWithQuantity
	if err := common.ReadJSON(r, &items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	input := pb.CreateOrderRequest{
		CustomerID: customerID,
		Items:      items,
	}

	order, err := h.grpcHandler.CreateOrderService(r.Context(), &input)
	if err != nil {
		common.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	common.WriteJSON(w, http.StatusCreated, order)
}
