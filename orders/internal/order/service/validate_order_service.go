package service

import (
	"context"

	common "github.com/charmingruby/remy-common"
	pb "github.com/charmingruby/remy-common/api"
	"github.com/google/uuid"
)

func (s *ServiceRegistry) ValidateOrderService(ctx context.Context, input *pb.CreateOrderRequest) ([]*pb.Item, error) {
	if len(input.Items) == 0 {
		return nil, common.NewNoItemsErr()
	}

	mergedItems := mergeItemsQuantities(input.Items)

	// temporary
	var itemsWithPrice []*pb.Item
	for _, i := range mergedItems {
		itemsWithPrice = append(itemsWithPrice, &pb.Item{
			Id:       uuid.NewString(),
			PriceId:  "price_1Pji1mGJRyNYkOQGwj2xSuZj",
			Quantity: i.Quantity,
		})
	}

	return itemsWithPrice, nil
}

func mergeItemsQuantities(items []*pb.ItemWithQuantity) []*pb.ItemWithQuantity {
	merged := []*pb.ItemWithQuantity{}

	for _, item := range items {
		var found bool = false

		for _, finalItem := range merged {
			if finalItem.Id == item.Id {
				finalItem.Quantity += item.Quantity
				found = true
				break
			}
		}

		if !found {
			merged = append(merged, item)
		}
	}

	return merged
}
