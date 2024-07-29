package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	_ "github.com/joho/godotenv/autoload"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	common "github.com/charmingruby/remy-common"
	"github.com/charmingruby/remy-gateway/internal/transport"
	"github.com/charmingruby/remy-gateway/internal/transport/grpc_transport"
	"github.com/charmingruby/remy-gateway/internal/transport/rest_transport"
)

var (
	httpAddr         = common.EnvString("HTTP_ADDR", ":8080")
	orderServiceAddr = common.EnvString("GRPC_ADDR", "localhost:2000")
)

func main() {
	conn, err := grpc.Dial(
		orderServiceAddr,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		log.Fatalf("Failed to dial gRPC server: %v", err)
	}
	defer conn.Close()
	log.Println("Dialing orders service at ", orderServiceAddr)

	grpcHandler := grpc_transport.NewHandler(conn)

	mux := http.NewServeMux()
	restHandler := rest_transport.NewHandler(mux, grpcHandler)
	restHandler.Register()

	srv := transport.NewHTTPServer(mux, httpAddr)
	log.Printf("Starting HTTP server at %s", httpAddr)
	if err := srv.Run(); err != nil {
		log.Fatal("Failed to start http server")
	}

	// gracefull shutdown
	term := make(chan os.Signal, 1)
	signal.Notify(term, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-term
		if err := srv.Server.Close(); !errors.Is(err, http.ErrServerClosed) {
			log.Fatalf("Error closing Server: %v", err)
		}
	}()
}
