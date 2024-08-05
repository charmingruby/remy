package gateway

import "context"

type OrdersGateway interface {
	UpdateOrdersAfterPaymentLink(ctx context.Context, orderID, paymentLink string) error
}
