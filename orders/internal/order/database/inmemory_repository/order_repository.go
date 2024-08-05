package inmemory_repository

import (
	"context"
	"errors"

	pb "github.com/charmingruby/remy-common/api"
	"github.com/google/uuid"
)

func NewOrderInMemoryRepository() *OrderInMemoryRepository {
	return &OrderInMemoryRepository{
		items: []*pb.Order{},
	}
}

type OrderInMemoryRepository struct {
	items []*pb.Order
}

func (r *OrderInMemoryRepository) Create(
	ctx context.Context,
	p *pb.CreateOrderRequest,
	items []*pb.Item) (string, error) {
	id := uuid.NewString()

	r.items = append(r.items, &pb.Order{
		ID:          id,
		CustomerID:  p.CustomerID,
		Status:      "pending",
		Items:       items,
		PaymentLink: "",
	})

	return id, nil
}

func (r *OrderInMemoryRepository) Get(ctx context.Context, id, customerID string) (*pb.Order, error) {
	for _, o := range r.items {
		if o.ID == id && o.CustomerID == customerID {
			return o, nil
		}
	}

	return nil, errors.New("order not found")
}

func (r *OrderInMemoryRepository) Update(ctx context.Context, id string, input *pb.Order) error {
	for i, o := range r.items {
		if o.ID == id {
			r.items[i].PaymentLink = input.PaymentLink
			r.items[i].Status = input.Status

			return nil
		}
	}

	return errors.New("order not found")
}
