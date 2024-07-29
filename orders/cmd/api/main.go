package main

import (
	"context"
	"log"

	common "github.com/charmingruby/remy-common"
	"github.com/charmingruby/remy-orders/internal/common/server"
	"github.com/charmingruby/remy-orders/internal/order"
	"github.com/charmingruby/remy-orders/internal/order/database/mongo_repository"
	"github.com/charmingruby/remy-orders/internal/order/domain"
)

var (
	grpcAddr = common.EnvString("GRPC_ADDR", "localhost:2000")
)

func main() {
	server := server.NewServer(grpcAddr)

	orderRepository := mongo_repository.NewOrderMongoRepository()
	orderService := domain.NewDomainServiceRegistry(orderRepository)
	orderService.CreateOrder(context.Background())
	_ = order.NewServiceRegistry(orderService)
	order.NewGRPCHandler(server.GRPCServer)

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
