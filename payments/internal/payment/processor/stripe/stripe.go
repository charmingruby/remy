package stripe

import (
	"fmt"
	"log"

	pb "github.com/charmingruby/remy-common/api"
	"github.com/stripe/stripe-go/v78"
	"github.com/stripe/stripe-go/v78/checkout/session"
)

func NewProcessor(gatewayAddr string) *Stripe {
	return &Stripe{
		GatewayAddr: gatewayAddr,
	}
}

type Stripe struct {
	GatewayAddr string
}

func (p *Stripe) CreatePaymentLink(order *pb.Order) (string, error) {
	log.Printf("Creating payment link for order %v", order)
	gatewaySuccessURL := fmt.Sprintf("%s/success.html?customerID=%s&orderID=%s",
		p.GatewayAddr,
		order.CustomerID,
		order.ID,
	)
	gatewayCancelURL := fmt.Sprintf("%s/cancel.html", p.GatewayAddr)

	items := []*stripe.CheckoutSessionLineItemParams{}
	for _, item := range order.Items {
		items = append(items, &stripe.CheckoutSessionLineItemParams{
			Price:    stripe.String(item.PriceId),
			Quantity: stripe.Int64(int64(item.Quantity)),
		})
	}

	params := &stripe.CheckoutSessionParams{
		LineItems:  items,
		Mode:       stripe.String(string(stripe.CheckoutSessionModePayment)),
		SuccessURL: stripe.String(gatewaySuccessURL),
		CancelURL:  stripe.String(gatewayCancelURL),
	}

	result, err := session.New(params)
	if err != nil {
		return "", err
	}

	return result.URL, nil
}
