package domain

import "context"

type OrderRepository interface {
	Create(context.Context) error
}
