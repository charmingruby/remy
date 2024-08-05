package rest

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"

	common "github.com/charmingruby/remy-common"
	pb "github.com/charmingruby/remy-common/api"
	"github.com/charmingruby/remy-common/broker"
	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/stripe/stripe-go"
	"github.com/stripe/stripe-go/webhook"
)

func (h *Handler) checkoutWebhookEndpoint(w http.ResponseWriter, r *http.Request) {
	const MaxBodyBytes = int64(65536)
	r.Body = http.MaxBytesReader(w, r.Body, MaxBodyBytes)

	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading request body: %v\n", err)
		w.WriteHeader(http.StatusServiceUnavailable)
		return
	}

	event, err := webhook.ConstructEvent(body, r.Header.Get("Stripe-Signature"), h.stripeEndpointSecret)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error verifying webhook signature: %v\n", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if event.Type == "payment_intent.succeeded" {
		var cs stripe.CheckoutSession
		err := json.Unmarshal(event.Data.Raw, &cs)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing webhook JSON: %v\n", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		order := &pb.Order{
			ID:          cs.Metadata["orderID"],
			CustomerID:  cs.Metadata["customerID"],
			Status:      "paid",
			PaymentLink: "",
		}

		jsonOrder, _ := json.Marshal(order)

		h.ch.PublishWithContext(ctx, broker.OrderPaidEvent, "", false, false, amqp.Publishing{
			ContentType:  "application/json",
			Body:         jsonOrder,
			DeliveryMode: amqp.Persistent,
		})

		fmt.Println("Message published order.paid")

		common.WriteJSON(w, http.StatusOK, body)
		return
	}

	w.WriteHeader(http.StatusOK)
}
