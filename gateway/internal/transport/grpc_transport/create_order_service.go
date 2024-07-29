package grpc_transport

import (
	"context"
	"errors"

	common "github.com/charmingruby/remy-common"
	pb "github.com/charmingruby/remy-common/api"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (h *Handler) CreateOrderService(
	ctx context.Context,
	req *pb.CreateOrderRequest,
) (*pb.Order, error) {
	order, err := h.Clients.OrderClient.CreateOrder(ctx, req)
	reqSts := status.Convert(err)
	if reqSts != nil {
		if reqSts.Code() != codes.InvalidArgument {
			return nil, common.NewPayloadErr(reqSts.Message())
		}

		return nil, errors.New(reqSts.Message())
	}

	return order, nil
}
