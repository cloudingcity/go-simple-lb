package server

import (
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

type Server struct {
	url     *url.URL
	handler http.Handler
}

func NewServer(u *url.URL) *Server {
	return &Server{
		url:     u,
		handler: httputil.NewSingleHostReverseProxy(u),
	}
}

func (s *Server) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	s.handler.ServeHTTP(rw, req)
}

func (s *Server) IsAlive() bool {
	_, err := net.DialTimeout("tcp", s.url.Host, 1*time.Second)
	return err == nil
}
