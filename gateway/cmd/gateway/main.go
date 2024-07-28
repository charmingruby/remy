package main

import (
	"log"
	"net/http"

	"github.com/charmingruby/remy-gateway/internal/transport"
	"github.com/charmingruby/remy-gateway/internal/transport/rest"
)

const (
	httpAddr = ":8000"
)

func main() {
	mux := http.NewServeMux()
	restHandler := rest.NewHandler(mux)
	restHandler.Register()

	server := transport.NewHTTPServer(mux, httpAddr)
	log.Printf("Starting HTTP server at %s", httpAddr)
	if err := server.Run(); err != nil {
		log.Fatal("Failed to start http server")
	}
}
