package server

import (
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
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

func (s *Server) IsAlive() bool {
	_, err := net.DialTimeout("tcp", s.URL.Host, 1*time.Second)
	return err == nil
}
