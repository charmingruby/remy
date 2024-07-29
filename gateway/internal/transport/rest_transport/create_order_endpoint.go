package rest_transport

import (
	"errors"
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

	if err := validateItems(items); err != nil {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	input := pb.CreateOrderRequest{
		CustomerID: customerID,
		Items:      items,
	}

	order, err := h.ordersGateway.CreateOrderService(r.Context(), &input)
	if err != nil {
		if isPayloadErr := errors.Is(err, common.NewPayloadErr(err.Error())); isPayloadErr {
			common.WriteError(w, http.StatusBadRequest, err.Error())
			return
		}

		common.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	common.WriteJSON(w, http.StatusCreated, order)
}

func validateItems(items []*pb.ItemWithQuantity) error {
	if len(items) == 0 {
		return errors.New("items must have at least one item")
	}

	for _, i := range items {
		if i.Id == "" {
			return errors.New("item ID is required")
		}

		if i.Quantity <= 0 {
			return errors.New("items must have a valid quantity")
		}
	}

	return nil
}
