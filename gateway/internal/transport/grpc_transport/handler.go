package grpc_transport

import (
	pb "github.com/charmingruby/remy-common/api"
	"google.golang.org/grpc"
)

func NewHandler(conn *grpc.ClientConn) *Handler {
	orderClient := pb.NewOrderServiceClient(conn)
	cls := ProtoClients{
		OrderClient: orderClient,
	}

	return &Handler{
		Conn:    conn,
		Clients: &cls,
	}
}

type Handler struct {
	Conn    *grpc.ClientConn
	Clients *ProtoClients
}

type ProtoClients struct {
	OrderClient pb.OrderServiceClient
}
