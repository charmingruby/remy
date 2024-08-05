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

func (g *gateway) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (*pb.Order, error) {
	conn, err := discovery.ServiceConnection(ctx, "orders", g.registry)
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}

	c := pb.NewOrderServiceClient(conn)

	input := pb.CreateOrderRequest{
		CustomerID: req.CustomerID,
		Items:      req.Items,
	}

	order, err := c.CreateOrder(ctx, &input)
	if err != nil {
		return nil, err
	}

	return order, nil
}
func (g *gateway) GetOrder(ctx context.Context, orderID, customerID string) (*pb.Order, error) {
	conn, err := discovery.ServiceConnection(ctx, "orders", g.registry)
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}

	c := pb.NewOrderServiceClient(conn)

	return c.GetOrder(ctx, &pb.GetOrderRequest{
		OrderID:    orderID,
		CustomerID: customerID,
	})
}
