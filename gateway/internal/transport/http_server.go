package transport

import (
	"net/http"
)

func NewHTTPServer(mux *http.ServeMux, addr string) *HTTPServer {
	return &HTTPServer{
		Server: &http.Server{
			Addr:    addr,
			Handler: mux,
		},
	}
}

type HTTPServer struct {
	Server *http.Server
}

func (s *HTTPServer) Run() error {
	if err := http.ListenAndServe(s.Server.Addr, s.Server.Handler); err != nil {
		return err
	}

	return nil
}
