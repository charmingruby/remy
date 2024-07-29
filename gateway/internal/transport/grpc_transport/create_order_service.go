package grpc_transport

import (
	"context"

	pb "github.com/charmingruby/remy-common/api"
)

func (h *Handler) CreateOrderService(
	ctx context.Context,
	req *pb.CreateOrderRequest,
) (*pb.Order, error) {
	return h.Clients.OrderClient.CreateOrder(ctx, req)
}
