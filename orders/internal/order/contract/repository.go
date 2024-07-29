package contract

import "context"

type OrderRepository interface {
	Create(context.Context) error
}
