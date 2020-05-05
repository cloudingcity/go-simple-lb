package server

import (
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type Server struct {
	serverURL *url.URL
	handler   http.Handler
}

func NewServer(u *url.URL) *Server {
	return &Server{
		serverURL: u,
		handler:   httputil.NewSingleHostReverseProxy(u),
	}
}

func (s *Server) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	s.handler.ServeHTTP(rw, req)
}

func (s *Server) IsAlive() bool {
	_, err := net.DialTimeout("tcp", s.serverURL.Host, 1*time.Second)
	return err == nil
}
