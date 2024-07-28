package main

import (
	"log"
	"net/http"

	_ "github.com/joho/godotenv/autoload"

	common "github.com/charmingruby/remy-common"
	"github.com/charmingruby/remy-gateway/internal/transport"
	"github.com/charmingruby/remy-gateway/internal/transport/rest"
)

func main() {
	httpAddr := common.EnvString("HTTP_ADDR", ":8080")

	mux := http.NewServeMux()
	restHandler := rest.NewHandler(mux)
	restHandler.Register()

	server := transport.NewHTTPServer(mux, httpAddr)
	log.Printf("Starting HTTP server at %s", httpAddr)
	if err := server.Run(); err != nil {
		log.Fatal("Failed to start http server")
	}
}
