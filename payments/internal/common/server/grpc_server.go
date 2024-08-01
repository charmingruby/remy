package server

import (
	"net"

	"google.golang.org/grpc"
)

func NewServer(addr string) *Server {
	server := grpc.NewServer()

	return &Server{
		GRPCServer: server,
		Addr:       addr,
	}
}

type Server struct {
	Addr       string
	GRPCServer *grpc.Server
}

func (s *Server) Run() (net.Listener, error) {
	l, err := net.Listen("tcp", s.Addr)
	if err != nil {
		return nil, err
	}

	return l, nil
}
