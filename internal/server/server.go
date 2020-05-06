package server

import (
	"net"
	"net/http"
	"net/url"
	"time"
)

type Server struct {
	url     *url.URL
	handler http.Handler
}

func NewServer(u *url.URL, handler http.Handler) *Server {
	return &Server{url: u, handler: handler}
}

func (s *Server) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	s.handler.ServeHTTP(rw, req)
}

func (s *Server) IsAlive() bool {
	_, err := net.DialTimeout("tcp", s.url.Host, 1*time.Second)
	return err == nil
}
