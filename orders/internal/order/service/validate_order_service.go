package service

import (
	"context"
	"log"

	common "github.com/charmingruby/remy-common"
	pb "github.com/charmingruby/remy-common/api"
)

func (s *ServiceRegistry) ValidateOrderService(ctx context.Context, input *pb.CreateOrderRequest) error {
	if len(input.Items) == 0 {
		return common.NewNoItemsErr()
	}

	mergedItems := mergeItemsQuantities(input.Items)
	log.Println(mergedItems)

	return nil
}

func mergeItemsQuantities(items []*pb.ItemWithQuantity) []*pb.ItemWithQuantity {
	merged := make([]*pb.ItemWithQuantity, 0)

	for _, item := range items {
		found := false

		for _, finalItem := range merged {
			if finalItem.Id == item.Id {
				finalItem.Quantity += item.Quantity
				found = true
				break
			}

			if !found {
				merged = append(merged, item)
			}
		}
	}

	return merged
}
