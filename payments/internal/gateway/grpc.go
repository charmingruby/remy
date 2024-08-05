package gateway

import (
	"context"
	"log"

	pb "github.com/charmingruby/remy-common/api"
	"github.com/charmingruby/remy-common/discovery"
)

type gateway struct {
	registry discovery.Registry
}

func NewGRPCGateway(registry discovery.Registry) *gateway {
	return &gateway{
		registry: registry,
	}
}

func (g *gateway) UpdateOrdersAfterPaymentLink(ctx context.Context, orderID, paymentLink string) error {
	conn, err := discovery.ServiceConnection(ctx, "orders", g.registry)
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}

	c := pb.NewOrderServiceClient(conn)

	input := &pb.Order{
		ID:          orderID,
		Status:      "waiting_payment",
		PaymentLink: paymentLink,
	}

	_, err = c.UpdateOrder(ctx, input)
	if err != nil {
		return err
	}

	return nil
}
