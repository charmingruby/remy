package order

import "github.com/charmingruby/remy-orders/internal/order/domain"

func NewServiceRegistry(svc domain.Service) *ServiceRegistry {
	return &ServiceRegistry{
		service: svc,
	}
}

type ServiceRegistry struct {
	service domain.Service
}
