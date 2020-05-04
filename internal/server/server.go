package server

import (
	"net/http"
	"net/http/httputil"
	"net/url"
)

type Server struct {
	URL     *url.URL
	handler http.Handler
}

func NewServer(serverURL *url.URL) *Server {
	return &Server{
		URL:     serverURL,
		handler: httputil.NewSingleHostReverseProxy(serverURL),
	}
}

func (s *Server) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	s.handler.ServeHTTP(rw, req)
}
