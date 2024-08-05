package rest

import (
	"errors"
	"fmt"
	"net/http"

	common "github.com/charmingruby/remy-common"
	pb "github.com/charmingruby/remy-common/api"
)

type CreaterOrderResponse struct {
	Order         *pb.Order `json:"order"`
	RedirectToURL string    `json:"redirect_to_url"`
}

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

	order, err := h.ordersGateway.CreateOrder(r.Context(), &input)
	if err != nil {
		if isPayloadErr := errors.Is(err, common.NewPayloadErr(err.Error())); isPayloadErr {
			common.WriteError(w, http.StatusBadRequest, err.Error())
			return
		}

		common.WriteError(w, http.StatusInternalServerError, err.Error())
		return
	}

	res := CreaterOrderResponse{
		Order:         order,
		RedirectToURL: fmt.Sprintf("http://localhost:8080/success.html?customerID=%s&orderID=%s", order.CustomerID, order.ID),
	}

	common.WriteJSON(w, http.StatusCreated, res)
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
