package transport

import "net/http"

func NewHTTPServer(mux *http.ServeMux, addr string) *HTTPServer {
	return &HTTPServer{
		mux:  mux,
		addr: addr,
	}
}

type HTTPServer struct {
	mux  *http.ServeMux
	addr string
}

func (s *HTTPServer) Run() error {
	if err := http.ListenAndServe(s.addr, s.mux); err != nil {
		return err
	}

	return nil
}
