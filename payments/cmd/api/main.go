package main

import (
	"context"
	"log"
	"time"

	_ "github.com/joho/godotenv/autoload"

	common "github.com/charmingruby/remy-common"
	"github.com/charmingruby/remy-common/broker"
	"github.com/charmingruby/remy-common/discovery"
	"github.com/charmingruby/remy-common/discovery/consul"
	"github.com/charmingruby/remy-payments/internal/common/server"
	"github.com/charmingruby/remy-payments/internal/payment"
	stripeProcessor "github.com/charmingruby/remy-payments/internal/payment/processor/stripe"
	"github.com/charmingruby/remy-payments/internal/payment/queue/rabbitmq"
	"github.com/stripe/stripe-go/v78"
)

var (
	serviceName  = "payment"
	grpcAddr     = common.EnvString("GRPC_ADDR", "localhost:2001")
	consulAddr   = common.EnvString("CONSUL_ADDR", "localhost:8500")
	amqpUser     = common.EnvString("RABBITMQ_USER", "guest")
	amqpPassword = common.EnvString("RABBITMQ_PASSWORD", "guest")
	amqpHost     = common.EnvString("RABBITMQ_HOST", "localhost")
	amqpPort     = common.EnvString("RABBITMQ_PORT", "5672")
	stripeAPIKey = common.EnvString("STRIPE_API_KEY", "")
	gatewayAddr  = common.EnvString("GATEWAY_ADDR", "http://localhost:8080")
)

func main() {
	// service discovery
	registry, err := consul.NewRegistry(consulAddr, serviceName)
	if err != nil {
		panic(err)
	}

	ctx := context.Background()
	instanceID := discovery.GenerateInstanceID(serviceName)
	if err := registry.Register(ctx, instanceID, serviceName, grpcAddr); err != nil {
		panic(err)
	}

	go func() {
		for {
			if err := registry.HealthCheck(instanceID, serviceName); err != nil {
				log.Fatal("Failed to health check")
			}

			time.Sleep(time.Second * 1)
		}
	}()
	defer registry.Deregister(ctx, instanceID, serviceName)

	// stripe setup
	stripe.Key = stripeAPIKey

	// broker conn
	ch, close := broker.Connect(amqpUser, amqpPassword, amqpHost, amqpPort)
	defer func() {
		close()
		ch.Close()
	}()

	paymentProcessor := stripeProcessor.NewProcessor(gatewayAddr)
	paymentSvc := payment.NewPaymentService(paymentProcessor)

	rabbitMQConsumer := rabbitmq.NewConsumer(paymentSvc)

	go rabbitMQConsumer.Listen(ch)

	// grpc server
	server := server.NewServer(grpcAddr)

	listener, err := server.Run()
	if err != nil {
		log.Fatalf("Failed to start gRPC server: %s", err.Error())
	}
	defer listener.Close()

	log.Printf("GRPC Server started at %s", grpcAddr)
	if err := server.GRPCServer.Serve(listener); err != nil {
		log.Fatal(err.Error())
	}
}
