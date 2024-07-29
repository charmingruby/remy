package grpc_transport

import (
	pb "github.com/charmingruby/remy-common/api"
	"google.golang.org/grpc"
)

func NewGRPCHandler(grpcServer *grpc.Server) {
	orderHandler := &gRPCOrderHandler{}
	pb.RegisterOrderServiceServer(grpcServer, orderHandler)
}
