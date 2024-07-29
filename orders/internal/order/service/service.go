package service

import (
	"github.com/charmingruby/remy-orders/internal/order/contract"
)

func NewServiceRegistry(orderRepository contract.OrderRepository) *ServiceRegistry {
	return &ServiceRegistry{
		orderRepository: orderRepository,
	}
}

type ServiceRegistry struct {
	orderRepository contract.OrderRepository
}
