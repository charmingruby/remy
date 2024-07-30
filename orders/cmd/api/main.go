package main

import (
	"context"
	"log"
	"time"

	common "github.com/charmingruby/remy-common"
	"github.com/charmingruby/remy-common/broker"
	"github.com/charmingruby/remy-common/discovery"
	"github.com/charmingruby/remy-common/discovery/consul"
	"github.com/charmingruby/remy-orders/internal/common/server"
	"github.com/charmingruby/remy-orders/internal/order"
	"github.com/charmingruby/remy-orders/internal/order/database/mongo_repository"
)

var (
	serviceName  = "orders"
	grpcAddr     = common.EnvString("GRPC_ADDR", "localhost:2000")
	consulAddr   = common.EnvString("CONSUL_ADDR", "localhost:8500")
	amqpUser     = common.EnvString("RABBITMQ_USER", "guest")
	amqpPassword = common.EnvString("RABBITMQ_PASSWORD", "guest")
	amqpHost     = common.EnvString("RABBITMQ_HOST", "localhost")
	amqpPort     = common.EnvString("RABBITMQ_PORT", "5672")
)

func main() {
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

	ch, close := broker.Connect(amqpUser, amqpPassword, amqpHost, amqpPort)
	defer func() {
		close()
		ch.Close()
	}()

	server := server.NewServer(grpcAddr)
	orderRepository := mongo_repository.NewOrderMongoRepository()
	orderService := order.NewServiceRegistry(orderRepository)
	order.NewGRPCHandler(server.GRPCServer, orderService, ch)

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
