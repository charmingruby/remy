package main

import (
	"context"

	"github.com/charmingruby/remy-orders/internal/order"
	"github.com/charmingruby/remy-orders/internal/order/database/mongo_repository"
	"github.com/charmingruby/remy-orders/internal/order/domain"
)

func main() {
	orderRepository := mongo_repository.NewOrderMongoRepository()
	orderService := domain.NewDomainServiceRegistry(orderRepository)
	orderService.CreateOrder(context.Background())
	_ = order.NewServiceRegistry(orderService)
}
