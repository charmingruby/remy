package main

import (
	"errors"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

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
