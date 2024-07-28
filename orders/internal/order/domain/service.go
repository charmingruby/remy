package domain

import "context"

type Service interface {
	CreateOrder(context.Context) error
}

func NewDomainServiceRegistry(orderRepository OrderRepository) *DomainServiceRegistry {
	return &DomainServiceRegistry{
		orderRepository: orderRepository,
	}
}

type DomainServiceRegistry struct {
	orderRepository OrderRepository
}

func (s *DomainServiceRegistry) CreateOrder(context.Context) error {
	return nil
}
