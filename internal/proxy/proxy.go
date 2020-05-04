package proxy

import (
	"fmt"
	"net/http"
	"net/url"

	"github.com/cloudingcity/go-simple-lb/internal/server"
	log "github.com/sirupsen/logrus"
)

type LoadBalancer struct {
	serverPool *server.Pool
}

func NewLB() *LoadBalancer {
	return &LoadBalancer{
		serverPool: server.NewPool(),
	}
}

func (lb *LoadBalancer) Add(u *url.URL) {
	lb.serverPool.Put(server.NewServer(u))
	log.Printf("Configured server: %s", u)
}

func (lb *LoadBalancer) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	lb.serverPool.GetNext().ServeHTTP(w, req)
}

func (lb *LoadBalancer) Listen(port int) {
	addr := fmt.Sprintf(":%d", port)
	log.Printf("Started listening on %s\n", addr)
	if err := http.ListenAndServe(addr, lb); err != nil {
		log.Fatal(err)
	}
}
